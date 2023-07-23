FROM golang:1.20.6-alpine3.17 AS build
LABEL maintainer="Rodrigo Prata <rbpsilver@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8000
CMD ["./main"]