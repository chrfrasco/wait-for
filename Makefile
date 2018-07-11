BINARY_NAME=bin/wait-for

all: test build

build:
	go build -o $(BINARY_NAME) -v
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go get ./...
