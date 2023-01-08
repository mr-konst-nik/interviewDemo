# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /demo ./cmd/demo/

EXPOSE 8080

CMD [ "/demo", "--config", "./config/cfg.yaml" ]