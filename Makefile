.PHONY: build

build:
	go build

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem
