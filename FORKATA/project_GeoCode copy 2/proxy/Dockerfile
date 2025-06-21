FROM golang:1.19-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o proxy ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/proxy    ./proxy
COPY --from=builder /app/migrations ./migrations

COPY --from=builder /app/docs     ./docs

EXPOSE 8080
CMD ["./proxy"]