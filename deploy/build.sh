#!/bin/bash
target=guser
mkdir -p ${target}
cd ..
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy/${target}/${target} cmd/app/main.go
go build -o deploy/${target}/${target} cmd/app/main.go
cd deploy