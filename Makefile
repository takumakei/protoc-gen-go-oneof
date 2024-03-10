.DEFAULT_GOAL := help

.PHONY: help example test

help: ## show this message
	@grep -E '^[.a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

example: ## $(MAKE) -C ./example gen test doc
	$(MAKE) -C ./example gen test doc

test: ## $(MAKE) -C ./test gen test
	$(MAKE) -C ./test gen test
