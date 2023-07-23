FROM golang:1.20.6-alpine3.17 AS build
LABEL maintainer="Rodrigo Prata <rbpsilver@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY docs ./docs
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo cmd/api/main.go

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=build /app/main .
RUN addgroup -S silver && adduser -S silver -G silver
USER silver
EXPOSE 8000
CMD ["./main"]