FROM golang:1.24.2-alpine3.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/main .

ENTRYPOINT ["/app/main"]
