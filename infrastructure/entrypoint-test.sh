#!/bin/sh

glide install
go test -v github.com/vehsamrak/telegramud/tests/...
