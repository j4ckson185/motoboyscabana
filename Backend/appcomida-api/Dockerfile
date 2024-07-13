FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY vendor ./vendor

COPY . .

RUN go build -o main ./cmd/api

EXPOSE 4003

CMD ["./main"]
