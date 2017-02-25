all: test build

build:
	go build -o bin/gotyouthere cmd

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
