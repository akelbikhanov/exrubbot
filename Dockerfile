# builder stage: используем свежий golang:alpine для компиляции
FROM golang:alpine AS builder

# Build arguments для версии
ARG VERSION=unknown
ARG GIT_COMMIT=unknown
ARG BUILD_TIME=unknown

# рабочая директория внутри контейнера-билдера
WORKDIR /app

# копируем модули отдельно, чтобы кэшировать зависимости
COPY go.mod go.sum ./
RUN go mod download

# копируем остальной исходный код
COPY . .

# компилируем статический бинарь для linux/amd64, убираем debug-символы
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
   go build \
   -ldflags="-s -w \
       -X 'github.com/akelbikhanov/exrubbot/internal/version.version=${VERSION}' \
       -X 'github.com/akelbikhanov/exrubbot/internal/version.gitCommit=${GIT_COMMIT}' \
       -X 'github.com/akelbikhanov/exrubbot/internal/version.buildTime=${BUILD_TIME}'" \
   -o /usr/local/bin/exrubbot ./cmd/bot

# stage для корневых сертификатов (Alpine минимального размера)
FROM alpine:latest AS certs
RUN apk add --no-cache ca-certificates

# финальный минимальный слой (scratch)
FROM scratch

# добавляем сертификаты для HTTPS-запросов
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

# добавляем скомпилированный бинарь в каталог, который уже в $PATH
COPY --from=builder /usr/local/bin/exrubbot /usr/local/bin/exrubbot

# запускаем под непривилегированным пользователем
USER 65532:65532

# точка входа контейнера
ENTRYPOINT ["exrubbot"]
