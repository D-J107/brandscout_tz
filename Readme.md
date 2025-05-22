
# Основной функционал

* Создание цитат 
* Получение всех цитат
* Получение цитат по автору
* Получение случайной цитаты

---

## Структура проекта

```plaintext
.
├── Dockerfile
├── Makefile
├── Readme.md
├── api
│   └── swagger.yaml
├── cmd
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   └── app
│       ├── config
│       │   └── config.go
│       ├── domain
│       │   └── models
│       │       └── quote.go
│       ├── repositories
│       │   ├── quote_dto.go
│       │   ├── quotes.go
│       │   └── quotes_simple_impl.go
│       ├── server.go
│       ├── services
│       │   ├── quotes.go
│       │   └── quotes_impl.go
│       ├── setup_routes.go
│       └── transport
│           └── rest
│               ├── error.go
│               ├── quotes_controller.go
│               └── quotes_dto.go
└── tests
    └── unit
        └── repo
            └── quotes_repo_test.go
```

---

## Запуск 
Склонируйте гит-репозиторий:
```bash
git clone https://github.com/D-J107/brandscout_tz.git
```
Запустите докер контейнер через compose:
```bash
docker compose up
```
Приложение станет доступно по порту :8080

---

## API схема

в файле api/swagger.yaml
Для демонстрации нужно зайти на любой сайт поддерживающий онлайн редактирование документа Swagger
Например: https://forge.etsi.org/swagger/editor/
При попытке запустить через "Try it out" выведется ошибка тк очевидно что сайты на удалённом хосте и не имеют доступа к моей локальной сети

---

## Тестирование
Запустите приложение.
Запустите команду вызова curl'ов через Makefile:
```bash
make run_curl_tests
```
она вызовет утилиту curl и сделает запросы:

```bash
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
{"id":"1","author":"Confucius","quote":"Life is simple, but we insist on making it complicated."}
curl http://localhost:8080/quotes
{"quotes":[{"author":"Confucius","quote":"Life is simple, but we insist on making it complicated."}]}
curl http://localhost:8080/quotes/random
{"author":"Confucius","quote":"Life is simple, but we insist on making it complicated."}
curl http://localhost:8080/quotes?author=Confucius
{"quotes":[{"author":"Confucius","quote":"Life is simple, but we insist on making it complicated."}]}
curl -X DELETE http://localhost:8080/quotes/1
```

curl отправляет запросы, получает корректные ответы.

* Запуск unit go тестов через команду
```bash
go test tests/unit/app/quotes_repo_test.go
```

---