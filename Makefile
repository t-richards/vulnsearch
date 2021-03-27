all: build

.PHONY: build
build: clean
	go build

.PHONY: clean
clean:
	rm -f coverage/*
	rm -f vulnsearch

.PHONY: deps
deps:
	go mod download
	go mod verify

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -coverprofile coverage/cover.out -v ./...
	go tool cover -html coverage/cover.out -o=coverage/index.html
