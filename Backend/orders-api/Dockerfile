FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Use o diretório vendor para dependências
RUN go build -mod=vendor -o main ./cmd/api

EXPOSE 4001

CMD ["./main"]
