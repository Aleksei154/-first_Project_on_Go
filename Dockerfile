# Используем официальный образ Go
FROM golang:1.24 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код приложения
COPY . .

# Собираем приложение статически
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app/main.go

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates

# Создаем директорию /app
RUN mkdir /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/main /app/

# Устанавливаем права на выполнение
RUN chmod +x /app/main

# Проверяем, что файл main существует
RUN ls -l /app

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]
