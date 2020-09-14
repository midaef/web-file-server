
.PHONY: build
build:
	go build -v ./cmd/webserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go run cmd/webserver/main.go -config-name config.json


.DEFAULT_GOAL := build