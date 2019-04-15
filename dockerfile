FROM golang:1.12 AS builder
RUN mkdir /orderfood
WORKDIR /orderfood
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /bin/orderfood

