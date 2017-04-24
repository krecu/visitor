#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o ./build/webserver-linux-x64 webserver.go
GOOS=linux GOARCH=386 go build -o ./build/webserver-linux-x64_86 webserver.go
GOOS=darwin GOARCH=amd64 go build -o ./build/webserver-darwin-x64 webserver.go
GOOS=darwin GOARCH=386 go build -o ./build/webserver-darwin-x64_86 webserver.go

GOOS=linux GOARCH=amd64 go build -o ./build/rpcserver-linux-x64 rpcserver.go
GOOS=linux GOARCH=386 go build -o ./build/rpcserver-linux-x64_86 rpcserver.go
GOOS=darwin GOARCH=amd64 go build -o ./build/rpcserver-darwin-x64 rpcserver.go
GOOS=darwin GOARCH=386 go build -o ./build/rpcserver-darwin-x64_86 rpcserver.go