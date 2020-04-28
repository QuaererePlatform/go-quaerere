
BINARY_NAME_COLUMBIA=columbia
BINARY_NAME_KOOTENAY=kootenay
BUILD_ROOT_COLUMBIA=cmd/columbia
BUILD_ROOT_KOOTENAY=cmd/kootenay
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
OUT_DIR=./out


all: test build

build: build_columbia build_kootenay

build_columbia:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(OUT_DIR)/$(BINARY_NAME_COLUMBIA) -v ./cmd/columbia

build_kootenay:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(OUT_DIR)/$(BINARY_NAME_KOOTENAY) -v ./cmd/kootenay

test:
	$(GO_TEST) -v ./...

clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME_COLUMBIA) $(BINARY_NAME_KOOTENAY)

deps:
	$(GO_GET)
