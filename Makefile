COMMIT 	= $(shell git log -1 --format=%H | cut -c1-8)
VERSION = 0.1.0

all: build

build: ## Build binary
	go build -ldflags '-X main.version=$(VERSION) -X main.commit=$(COMMIT)'

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
