#!/bin/bash
source ~/.bash_profile
go mod tidy
go run /app/cmd/gRPC/server/main.go
