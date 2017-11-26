FROM golang:1.9-alpine

WORKDIR /go/src/github.com/vehsamrak/telegramud
ADD . .

RUN apk add --no-cache git curl
RUN curl https://glide.sh/get | sh
