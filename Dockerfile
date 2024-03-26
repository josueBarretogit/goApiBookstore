# syntax=docker/dockerfile:1

FROM golang:1.21.1

WORKDIR /app

COPY . ./

RUN go mod download

RUN  go build -o /main

CMD ["/main"]
