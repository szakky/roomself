FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o todo-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/todo-api .

COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./todo-api"]