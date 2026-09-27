# daysNewYear

Сервис на Go, вычисляющий количество календарных дней, оставшихся до ближайшего 1 января.

## Требования

Go версии 1.22 или новее

## Сборка

    go build ./...

## Запуск

    go run ./cmd/server

Сервер стартует на порту 8080.

## Тесты

    go test ./... -v

## Docker

    docker build -t daysnewyear .
    docker run --rm -p 8080:8080 daysnewyear

## HTTP API

## GET /days-left

Возвращает количество дней до ближайшего 1 января.

Пример запроса без даты:

    curl "http://localhost:8080/days-left"

Пример запроса с датой:

    curl "http://localhost:8080/days-left?date=2026-12-25"

Успешный ответ (200 OK):

    {"date": "2026-12-25", "days_left": 7}

Ответ при некорректном формате даты (400 Bad Request):

    {"error": "неверный формат даты, ожидается YYYY-MM-DD"}

## GET /healthz

Возвращает статус сервиса в формате JSON:

    {"status":"ok"}

JSON выбран для машиночитаемости и единообразия с основным API.

## Benchmark

    go test -bench=. -benchmem ./internal/daysleft/

## Покрытие тестами

    go test -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

## Структура проекта

    cmd/server/          — точка входа HTTP-сервера
    internal/daysleft/   — вычислительная логика (не зависит от HTTP)
    internal/httpapi/    — HTTP-обработчики и middleware
    .github/workflows/   — конфигурация CI (GitHub Actions)

## CI

При каждом push и pull request в ветку main автоматически выполняются: форматирование (gofmt), статический анализ (go vet), проверка актуальности go.mod и go.sum (go mod tidy), проверка зависимостей (go mod verify), сборка (go build), тесты с детектором гонок и покрытием (go test -race -coverprofile), отчёт покрытия, загрузка артефакта покрытия и проверка уязвимостей (govulncheck).