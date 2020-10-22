# todo-list

Пример приложения тудушки из обучалки со goswagger.io

## Требования

- [goswagger](https://goswagger.io/ "Swagger")
- [golangci-lint](https://golangci-lint.run/ "Линтер")

## todo-list-server

Пример сервера с моделями и апишкой сгенерированной из свагера

## todo-list-client

Пример консольного клиента для взаимодействия с сервером

## Основные команды Makefile

### Собрать исполняемые файлы
Производит 2 файла: server и client
``` shell
make
```

### Проверка
Выполняет проверку набором линтеров и проверку спецификации
``` shell
make check
```

### Генерация кода из спецификации Swagger
Создаёт код необходимый клиенту и серверу
``` shell
make generate
```

### Развернуть веб-интерфейс с документацией по интерфейсу
По умолчанию запускает Swagger-UI
``` shell
make ui
```
Если указать параметр FLAVOR=redoc, то запустится Redoc-UI
``` shell
make ui FLAVOR=redoc
```
