# URL Shortener

Сервис сокращения ссылок на Go с PostgreSQL и Docker.

## Возможности

- `POST /shorten` создаёт короткую ссылку
- `GET /{code}` перенаправляет на оригинальный адрес (302)
- `GET /stats/{code}` показывает количество переходов
- Хранилище: PostgreSQL (или память, если `DATABASE_URL` не задан)

## Запуск

```bash
cp .env.example .env
docker compose up --build
```

## Примеры

```bash
curl -X POST localhost:8080/shorten -d '{"url":"https://go.dev"}'
curl -i localhost:8080/КОД
curl localhost:8080/stats/КОД
```

## Тесты

```bash
go test ./...
```

## Структура

- `cmd/server` — точка входа
- `internal/handler` — HTTP-хендлеры
- `internal/storage` — хранилища (память и PostgreSQL)
- `internal/config` — конфиг из переменных окружения
- `migrations` — SQL-миграции