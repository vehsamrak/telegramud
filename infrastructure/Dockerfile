FROM golang:1.9-alpine

WORKDIR /go/src/github.com/vehsamrak/telegramud

RUN apk add --no-cache git curl
RUN curl https://glide.sh/get | sh
RUN go get github.com/stretchr/testify
RUN go get github.com/vehsamrak/osinterrupt
