#!/usr/bin/env bash
set -xe

# install packages and dependencies
go get -u github.com/tools/godep
go get -u github.com/stretchr/testify
go get -u golang.org/x/lint/golint

# build command
go build -o bin/application application.go

