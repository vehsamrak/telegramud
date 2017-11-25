FROM golang:1.9-alpine

WORKDIR /go/src/github.com/vehsamrak/telegramud

RUN apk add --no-cache git
RUN glide install
RUN go get github.com/stretchr/testify
RUN go get github.com/vehsamrak/osinterrupt
