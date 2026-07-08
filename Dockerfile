# Собираем Go-приложение
FROM golang:latest AS builder

# Рабочая папка для сборки
WORKDIR /src

# Копируем зависимости
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем Linux-исполняемый файл
RUN CGO_ENABLED=0 GOOS=linux go build -o /book-api ./cmd/api

# Финальный лёгкий контейнер
FROM alpine:latest

# Рабочая папка приложения
WORKDIR /app

# Копируем собранный файл
COPY --from=builder /book-api /app/book-api

# Создаём папку для SQLite
RUN mkdir -p /app/data

# Порт приложения
EXPOSE 8080

# Запуск приложения
CMD ["/app/book-api"]