#!/usr/bin/env bash
export PATH=$PATH:$GOPATH/protobuf/bin
go build main.go && go run main.go