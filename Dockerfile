# Этап 1: Сборка приложения
FROM golang:1.24 AS builder


WORKDIR /app

# Кэширование зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o exrubbot ./cmd/bot/

# Этап 2: Минимальный контейнер для запуска
FROM alpine:latest

# Установка SSL-сертификатов для https (очень важно для телеграмм бота!)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем только бинарник из предыдущего этапа
COPY --from=builder /app/exrubbot .

# Запуск приложения
CMD ["./exrubbot"]
