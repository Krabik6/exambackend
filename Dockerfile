## Базовый образ для установки зависимостей и копирования исходного кода
#FROM golang:1.21 AS builder
#
#WORKDIR /app
#COPY . .
#
## Установка зависимостей и сборка исполняемого файла
#RUN go mod download
#RUN go build -o app ./cmd/main.go
#
## Образ для выполнения приложения
#FROM alpine AS runtime
#WORKDIR /root/
#
## Установка зависимостей для работы приложения
#RUN apk --no-cache add ca-certificates
#
## Копирование исполняемого файла из этапа builder
#COPY --from=builder /app/app .
#
## Задайте команду для выполнения приложения
#CMD ["./app"]


# Базовый образ для установки зависимостей и копирования исходного кода
FROM golang:1.21 AS base


# Установите переменную окружения GOPATH
ENV GOPATH=/

# Копируйте исходный код
COPY ./ /app

# Соберите Go-приложение
WORKDIR /app
RUN go mod download
RUN go build -o damn ./cmd/main.go

# Образ для выполнения приложения
FROM base AS damn

# Задайте команду для выполнения приложения
CMD ["./damn"]

# Образ для выполнения тестов
FROM base AS test

# Задайте команду для выполнения тестов
CMD ["go", "test", "./..."]
