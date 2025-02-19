.PHONY: build run clean test fmt deps

BINARY_NAME=todolist-server

deps:
	go mod download
	go mod tidy

build: deps
	go build -o $(BINARY_NAME) main.go

run: deps
	go run main.go

clean:
	rm -f $(BINARY_NAME)
	go clean

test: deps
	go test ./...

fmt:
	go fmt ./...

# 默认目标
default: build