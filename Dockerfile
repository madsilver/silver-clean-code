FROM golang:1.20.6-alpine3.17 AS build

LABEL maintainer="Rodrigo Prata <rbpsilver@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/api/main.go

EXPOSE 8000

CMD ["./main"]