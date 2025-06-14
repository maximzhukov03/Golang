# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Копируем только необходимое
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]