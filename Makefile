all: build

.PHONY: assets
assets:
	cd public && tsc

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

.PHONY: generate
generate:
	go generate

.PHONY: lint
lint:
	go vet ./...

.PHONY: start
start: clean
	go run -race -tags=dev main.go

.PHONY: test
test:
	go test -race -coverprofile coverage/cover.out -v ./...
	go tool cover -html coverage/cover.out -o=coverage/index.html
