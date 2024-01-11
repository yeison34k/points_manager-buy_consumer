.PHONY: build

test:
	go test ./... -v

build:
	sam build
