# Название приложения
APPNAME = TodoList

# Названия бинарей
BINS = client server

SWAGGER_SPEC = api/swagger.yaml

# URL реестра докерных образов
REGISTRY ?= docker.pkg.github.com/axard/todo-list

# Версия сборки из последнего тэга
VERSION := $(shell git describe --tags --always --dirty)

# Папки проекта с исходниками
SRC_DIRS := cmd internal

# Платформы для которых производить сборку
ALL_PLATFORMS := linux/amd64 windows/amd64

# Операционная система и архитектура
#   NOTE:
#   получение с помощью утилиты go накладывает необходимость её наличия в системе
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

GOPATH := $(if $(GOPATH),$(GOPATH),$(shell go env GOPATH))

BASEIMAGE ?= gcr.io/distroless/static

BUILD_IMAGE ?= golang:1.14-alpine
SWAGGER_IMAGE ?= quay.io/goswagger/swagger
GOLANGCI_IMAGE ?= golangci/golangci-lint:v1.31.0

BIN_EXTENSION :=
ifeq ($(OS), windows)
  BIN_EXTENSION := .exe
endif

all: # @HELP Сборка бинарей для одной платформы ($OS/$ARCH)
all: build

build-%:
	@$(MAKE) build                        \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

container-%:
	@$(MAKE) container                    \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

push-%:
	@$(MAKE) push                         \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

all-build: # @HELP Собирает бинарники для всех платформ
all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))

all-container: # @HELP Собирает контейнеры для всех платформ
all-container: $(addprefix container-, $(subst /,_, $(ALL_PLATFORMS)))

all-push: # @HELP Отправляет контейнеры для всех платформ в реестр
all-push: $(addprefix push-, $(subst /,_, $(ALL_PLATFORMS)))

# Папки куда будут складываться бинарники
OUTBINS = $(foreach bin,$(BINS),bin/$(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION))

# Сборка
build: $(OUTBINS)

# Папки необходимые для сборки и тестов
BUILD_DIRS := bin/$(OS)_$(ARCH)     \
              .go/bin/$(OS)_$(ARCH) \
              .go/cache

# Это хак позволяющий повзоляющий обойти гошное поведение постоянно изменяющее
# временную метку файла
# Каждый бинарь - это фасад для соответствующего отпечатка
# Создаём прямую связь между бинарём и отпечатком
$(foreach outbin,$(OUTBINS),$(eval  \
    $(outbin): .go/$(outbin).stamp  \
))

$(OUTBINS):
	@true

# Каждый бинарь - это фасад для соответствующего отпечатка
# Создаём обратную связь между бинарём и отпечатком
$(foreach outbin,$(OUTBINS),$(eval $(strip   \
    .go/$(outbin).stamp: OUTBIN = $(outbin)  \
)))

# Запускаем сборку для всех бинарей в папке ./.go и обновляем реальные бинари,
# если надо изменяем бинарник в ./bin
STAMPS = $(foreach outbin,$(OUTBINS),.go/$(outbin).stamp)
.PHONY: $(STAMPS)
$(STAMPS): go-build
	@echo "binary: $(OUTBIN)"
	@if ! cmp -s .go/$(OUTBIN) $(OUTBIN); then  \
	    mv .go/$(OUTBIN) $(OUTBIN);             \
	    date >$@;                               \
	fi

# Настоящие действия по сборке
go-build: $(BUILD_DIRS)
	@echo
	@echo "building for $(OS)/$(ARCH)"
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh -c "                                            \
	        ARCH=$(ARCH)                                        \
	        OS=$(OS)                                            \
	        VERSION=$(VERSION)                                  \
	        ./scripts/build.sh                                  \
	    "

# Пример: make shell CMD="-c 'date > datefile'"
shell: # @HELP Запускает командную оболочку внутри контейнера
shell: $(BUILD_DIRS)
	@echo "launching a shell in the containerized build environment"
	@docker run                                                 \
	    -ti                                                     \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh $(CMD)

CONTAINER_DOTFILES = $(foreach bin,$(BINS),.container-$(subst /,_,$(REGISTRY)/$(bin))-$(TAG))

container containers: # @HELP Собирает контейнер(ы) для каждого бинаря ($OS/$ARCH)
container containers: $(CONTAINER_DOTFILES)
	@for bin in $(BINS); do              \
	    echo "container: $(REGISTRY)/$$bin:$(TAG)"; \
	done

# Каждый дотфайл контейнерной цели может ссылаться на бинарь
# Сделано в 2 шага чтобы использовать специфичные для цели переменные
$(foreach bin,$(BINS),$(eval $(strip                                 \
    .container-$(subst /,_,$(REGISTRY)/$(bin))-$(TAG): BIN = $(bin)  \
)))
$(foreach bin,$(BINS),$(eval                                                                   \
    .container-$(subst /,_,$(REGISTRY)/$(bin))-$(TAG): bin/$(OS)_$(ARCH)/$(bin) build/Dockerfile.in  \
))
# Определение цели для всех дотфайлов контейнеров
$(CONTAINER_DOTFILES):
	@sed                                 \
	    -e 's|{ARG_BIN}|$(BIN)|g'        \
	    -e 's|{ARG_ARCH}|$(ARCH)|g'      \
	    -e 's|{ARG_OS}|$(OS)|g'          \
	    -e 's|{ARG_FROM}|$(BASEIMAGE)|g' \
	    build/Dockerfile.in > .dockerfile-$(BIN)-$(OS)_$(ARCH)
	@docker build -t $(REGISTRY)/$(BIN):$(TAG) -f .dockerfile-$(BIN)-$(OS)_$(ARCH) .
	@docker images -q $(REGISTRY)/$(BIN):$(TAG) > $@
	@echo

push: # @HELP Отправить контейнер с архитектурой ($OS/$ARCH) в ранее определённый реестр образов
push: $(CONTAINER_DOTFILES)
	@for bin in $(BINS); do                    \
	    docker push $(REGISTRY)/$$bin:$(TAG);  \
	done

manifest-list: # @HELP создаёт манифест списка контейнеров для всех платформ
manifest-list: all-push
	@for bin in $(BINS); do                                   \
	    platforms=$$(echo $(ALL_PLATFORMS) | sed 's/ /,/g');  \
	    manifest-tool                                         \
	        --username=oauth2accesstoken                      \
	        --password=$$(gcloud auth print-access-token)     \
	        push from-args                                    \
	        --platforms "$$platforms"                         \
	        --template $(REGISTRY)/$$bin:$(VERSION)__OS_ARCH  \
	        --target $(REGISTRY)/$$bin:$(VERSION)

version: # @HELP Выводит версию ПО
version:
	@echo $(VERSION)

test: # @HELP Запускает тесты с помощью скрипта ./scripts/test.sh
test: $(BUILD_DIRS)
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/sh -c "                                            \
	        ARCH=$(ARCH)                                        \
	        OS=$(OS)                                            \
	        VERSION=$(VERSION)                                  \
	        ./scripts/test.sh $(SRC_DIRS)                         \
	    "

check-code: # @HELP Проверяет код линтером
check-code:
	@echo
	@echo "linting golang code"
	@docker run                                                 \
	    --rm                                                    \
	    -v $$(pwd):/app                                         \
	    -w /app                                                 \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(GOLANGCI_IMAGE)                                       \
			golangci-lint run                                       \
			--color always

check-spec: # @HELP Проверяет Swagger (OpenAPI 2.0) документацию
check-spec:
	@echo
	@echo "validating Swagger spec"
	@docker run                                                 \
	    -it                                                     \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/go/src                                      \
	    -w /go/src                                              \
			--env GOPATH=$(GOPATH):/go                              \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(SWAGGER_IMAGE)                                        \
			validate $(SWAGGER_SPEC)

FLAVOR ?= swagger
show: # @HELP Показывает страничку Swagger-UI с документацией в представлении ($FLAVOR)
show:
	@echo
	@echo "showing docs with Swagger UI"
	@docker run                                                 \
	    -it                                                     \
	    --rm                                                    \
			-p 8888:8888                                            \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/go/src                                      \
	    -w /go/src                                              \
			--env GOPATH=$(GOPATH):/go                              \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(SWAGGER_IMAGE)                                        \
			serve $(SWAGGER_SPEC)                                   \
      --flavor=$(FLAVOR)                                      \
      --port=8888                                             \
      --no-open

boilerplate-server: # @HELP Генерирует код сервера API по спецификации OpenAPI 2.0
boilerplate-server:
	@echo
	@echo "generating server code from openapi specification"
	docker run                                                  \
	    -it                                                     \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $(HOME):$(HOME)                                      \
	    -w $$(pwd)                                              \
			--env GOPATH=$(GOPATH):/go                              \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(SWAGGER_IMAGE)                                        \
			generate server                                         \
			--exclude-main                                          \
      --model-package=restmodels                              \
      -f $(SWAGGER_SPEC)                                      \
      -t internal                                             \
      -A $(APPNAME)

boilerplate-client: # @HELP Генерирует код сервера API по спецификации OpenAPI 2.0
boilerplate-client:
	@echo
	@echo "generating client code from openapi specification"
	@docker run                                                 \
	    -it                                                     \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $(HOME):$(HOME)                                      \
	    -w $$(pwd)                                              \
			--env GOPATH=$(GOPATH):/go                              \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(SWAGGER_IMAGE)                                        \
			generate client                                         \
			--skip-models                                           \
      --model-package=restmodels                              \
      -f $(SWAGGER_SPEC)                                      \
      -t internal                                             \
      -A $(APPNAME)

$(BUILD_DIRS):
	@mkdir -p $@

clean: # @HELP Удаляет собранные контейнеры и временные файлы
clean: container-clean bin-clean

container-clean:
	rm -rf .container-* .dockerfile-*

bin-clean:
	rm -rf .go bin

help: # @HELP Выводит эту справку
help:
	@echo "VARIABLES:"
	@echo "  BINS = $(BINS)"
	@echo "  OS = $(OS)"
	@echo "  ARCH = $(ARCH)"
	@echo "  REGISTRY = $(REGISTRY)"
	@echo
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)    \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '
