.DEFAULT_GOAL := all

.PHONY: vulnsearch
vulnsearch: ## Compile application
	crystal build -o bin/vulnsearch src/vulnsearch.cr

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: dist
dist: ## Compile application for production
	crystal build --release -o bin/vulnsearch src/vulnsearch.cr

.PHONY: dev
dev: ## Run application in development
	crystal run src/vulnsearch.cr

.PHONY: clean
clean: ## Remove compiled binaries and downloaded dependencies
	rm -f bin/vulnsearch
	rm -rf .shards
	rm -rf lib

.PHONY: deps
deps: ## Install required Crystal dependencies via shards
	shards install

.PHONY: db
db: ## Migrate database
	bin/micrate up

.PHONY: test
test: ## Run unit tests / specs
	KEMAL_ENV=test crystal spec spec/vulnsearch.cr

.PHONY: all ## Targets required to run application
all:
	$(MAKE) deps
	$(MAKE) db
	$(MAKE) vulnsearch
