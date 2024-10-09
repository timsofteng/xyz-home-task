ARG APP_NAME
ARG HTTP_SERVER_PORT
FROM golang:1.23.0-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go test ./...
RUN CGO_ENABLED=0 go build -o bin/$APP_NAME ./cmd/*

FROM alpine:latest

WORKDIR /
COPY --from=build-stage /app/bin/$APP_NAME app/$APP_NAME

EXPOSE $HTTP_SERVER_PORT

USER nobody
CMD ./app/*

