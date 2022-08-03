#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o guser cmd/app/main.go
