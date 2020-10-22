MODULE = github.com/axard/todo-list

OUT = out

.PHONY = clean, tidy, check, install

build: $(OUT)/todo-list-server $(OUT)/todo-list-client

$(OUT)/todo-list-server: cmd/todo-list-server/main.go | $(OUT)/
	@echo "+----------------------------+"
	@echo "| Сборка сервера «todo-list» |"
	@echo "+----------------------------+"
	go build -o $@ ./cmd/todo-list-server

$(OUT)/todo-list-client: cmd/todo-list-client/main.go | $(OUT)/
	@echo "+----------------------------+"
	@echo "| Сборка клиента «todo-list» |"
	@echo "+----------------------------+"
	go build -o $@ ./cmd/todo-list-client

$(OUT)/:
	@mkdir -p $@

tidy:
	@echo "+----------------------+"
	@echo "| Очистка зависимостей |"
	@echo "+----------------------+"
	go mod tidy

clean:
	-rm -rf $(OUT)/

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
