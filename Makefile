# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
BINARY_NAME=kootenay

all: test build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BINARY_NAME) -v

test:
	$(GO_TEST) -v ./...

clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

run:
	./$(BINARY_NAME)

deps:
	$(GO_GET)
