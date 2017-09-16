.DEFAULT_GOAL := all

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: vulnsearch
vulnsearch: ## Compile application
	shards build

.PHONY: dist
dist: ## Compile application for production
	shards build --release
	strip bin/vulnsearch

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
	$(MAKE) dist
	$(MAKE) db
