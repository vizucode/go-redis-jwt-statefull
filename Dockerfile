FROM golang:alpine

RUN apk update && apk add --no-cache git

ENV SERVER_PORT = "8080"

ENV DB_HOST = "localhost"
ENV DB_USER = "root"
ENV DB_PASSWORD = "vizu"
ENV DB_NAME = "test"
ENV DB_PORT = "5432"

ENV REDIS_ADDR = "6379"
ENV REDIS_USERNAME = "havis"
ENV REDIS_PASSWORD = "secret"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]