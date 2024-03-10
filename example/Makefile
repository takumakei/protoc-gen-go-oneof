.DEFAULT_GOAL := help

.PHONY: help gen doc test

help: ## show this message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-8s\033[0m %s\n", $$1, $$2}'

gen: ## buf generate
	rm -fr gen && \
	buf format -w && \
	buf lint && \
	buf generate

test: ## go test -v ./...
	go test -v ./...

doc: ## go doc -all ./gen/example/v1
	go doc -all ./gen/example/v1
