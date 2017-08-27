FROM golang:1.8-alpine

WORKDIR /go/src/github.com/Vehsamrak/telegramud
COPY . .

RUN apk add --no-cache git
RUN go get -v ./internal/...
RUN go get github.com/stretchr/testify
RUN go test -v ./tests/...
RUN go build -o build/telegramud
