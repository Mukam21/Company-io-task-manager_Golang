FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o io-bound-task-api ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/io-bound-task-api .

EXPOSE 8080

CMD ["./io-bound-task-api"]
