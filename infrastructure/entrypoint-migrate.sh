#!/bin/sh

glide install
go run infrastructure/migrations/migrate.go
