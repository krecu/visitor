#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o ./build/visitor-linux-x64 main.go
GOOS=linux GOARCH=386 go build -o ./build/visitor-linux-x64_86 main.go
GOOS=darwin GOARCH=amd64 go build -o ./build/visitor-darwin-x64 main.go
GOOS=darwin GOARCH=386 go build -o ./build/visitor-darwin-x64_86 main.go