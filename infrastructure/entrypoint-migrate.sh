#!/bin/sh

go run infrastructure/migrations/migrate.go
go run github.com/vehsamrak/telegramud/infrastructure/migrations/migrate.go
go run /go/src/github.com/vehsamrak/telegramud/infrastructure/migrations/migrate.go
