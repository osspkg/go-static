SHELL=/bin/bash

install:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/static ./cmd/static/

tests:
	go test -race -v ./...