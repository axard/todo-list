MODULE = github.com/axard/todo-list

COMMON_DIRS = internal/store internal/restmodels
COMMON_FILES = $(shell find $(COMMON_DIRS) -type f -name "*.go")

CLIENT_DIRS = internal/client cmd/todo-list-client
CLIENT_FILES = $(shell find $(CLIENT_DIRS) -type f -name '*.go')

SERVER_DIRS = internal/restapi cmd/todo-list-server
SERVER_FILES = $(shell find $(SERVER_DIRS) -type f -name '*.go')

.PHONY = build, server, client, clean, tidy, check, install

build: server client

server: $(COMMON_FILES) $(SERVER_FILES)
	@echo "+----------------------------+"
	@echo "| Сборка сервера «todo-list» |"
	@echo "+----------------------------+"
	go build -o server ./cmd/todo-list-server

client: $(COMMON_FILES) $(CLIENT_FILES)
	@echo "+----------------------------+"
	@echo "| Сборка клиента «todo-list» |"
	@echo "+----------------------------+"
	go build -o client ./cmd/todo-list-client

tidy:
	@echo "+----------------------+"
	@echo "| Очистка зависимостей |"
	@echo "+----------------------+"
	go mod tidy

clean:
	-rm -rf server
	-rm -rf client

check: check-golang-code check-swagger-spec

check-golang-code:
	@echo "+---------------+"
	@echo "| Проверка кода |"
	@echo "+---------------+"
	golangci-lint run

check-swagger-spec:
	@echo "+---------------------------+"
	@echo "| Проверка спецификации API |"
	@echo "+---------------------------+"
	swagger validate api/swagger.yaml

install: install-server install-client

install-server:
	@echo "+-------------------------------+"
	@echo "| Установка сервера «todo-list» |"
	@echo "+-------------------------------+"
	go install ./cmd/todo-list-server

install-client:
	@echo "+-------------------------------+"
	@echo "| Установка клиента «todo-list» |"
	@echo "+-------------------------------+"
	go install ./cmd/todo-list-client

generate: generate-server generate-client

generate-server:
	@echo "+--------------------------------------------+"
	@echo "| Генерация кода сервера по спецификации API |"
	@echo "+--------------------------------------------+"
	swagger generate server --exclude-main --model-package=restmodels -f api/swagger.yaml -t internal -A TodoList

generate-client:
	@echo "+--------------------------------------------+"
	@echo "| Генерация кода клиента по спецификации API |"
	@echo "+--------------------------------------------+"
	swagger generate client --skip-models --model-package=restmodels -f api/swagger.yaml -t internal -A TodoList

FLAVOR = swagger

ui:
	swagger serve ./api/swagger.yaml --flavor=$(FLAVOR)
