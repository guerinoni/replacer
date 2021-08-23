.PHONY: build

build:
	go build

test:
	go test -race -shuffle on ./...

bench:
	go test -bench=. -benchmem
