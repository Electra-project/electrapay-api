NAME = electrapay-api

ifeq ($(OS), Windows_NT)
	BINARY_NAME = ${NAME}.exe
else
	BINARY_NAME = ${NAME}
endif

install:
	go get -u github.com/tools/godep
	go get -u github.com/stretchr/testify
	go get -u golang.org/x/lint/golint
	godep restore

lint:
	golint ./application.go
	golint ./src/...

start:
	go build -ldflags="-s -w" && "./electrapay-api"

test:
	make lint
	go test -v ./...
