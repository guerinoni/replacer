.PHONY: build

build:
	go build

test:
	go test -v -race ./...

gofmt: ## Run gofmt locally without overwriting any file
	gofmt -d -s $$(find . -name '*.go' | grep -v vendor)

govet: ## Run go vet on the project
	go vet ./...

