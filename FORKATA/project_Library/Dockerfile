FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/library cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/library .
COPY .env .

EXPOSE 8080

CMD ["./library"]