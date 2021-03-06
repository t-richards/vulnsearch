all: build

.PHONY: assets
assets:
	cd src && npm run build

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
generate: assets
	go generate

.PHONY: lint
lint:
	go vet ./...

.PHONY: npm
npm:
	cd src && npm install

.PHONY: start
start: clean
	go run -race -tags=dev main.go

.PHONY: test
test:
	go test -race -coverprofile coverage/cover.out -v ./...
	go tool cover -html coverage/cover.out -o=coverage/index.html
