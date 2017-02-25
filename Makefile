all: test build vet

build:
	go build -o bin/gotyouthere cmd/**.go

vet:
	go vet ./...

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic cmd/*.go
	go test -race -coverprofile=coverage.txt -covermode=atomic pkg/server/*.go
