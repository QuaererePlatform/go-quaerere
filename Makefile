.PHONY: all build build_columbia build_kootenay proto test clean deps db_init_kootenay

BINARY_NAME_COLUMBIA=columbia
BINARY_NAME_KOOTENAY=kootenay
BUILD_ROOT_COLUMBIA=cmd/columbia
BUILD_ROOT_KOOTENAY=cmd/kootenay
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_MOD=$(GO_CMD) mod
OUT_DIR=./out
PROTOC=protoc
PROTO_PATH=api/proto/v0


all: proto test build

build: build_columbia build_kootenay

build_columbia:
	CGO_ENABLED=0 $(GO_BUILD) -o $(OUT_DIR)/$(BINARY_NAME_COLUMBIA) -v ./cmd/columbia

build_kootenay:
	CGO_ENABLED=0 $(GO_BUILD) -o $(OUT_DIR)/$(BINARY_NAME_KOOTENAY) -v ./cmd/kootenay

proto:
	$(PROTOC) --proto_path=$(PROTO_PATH) --proto_path=third_party --go_out=plugins=grpc:. $(PROTO_PATH)/accounting.proto
	$(PROTOC) --proto_path=$(PROTO_PATH) --proto_path=third_party --go_out=plugins=grpc:. $(PROTO_PATH)/web_page.proto
	$(PROTOC) --proto_path=$(PROTO_PATH) --proto_path=third_party --go_out=plugins=grpc:. $(PROTO_PATH)/web_site.proto

test: proto
	$(GO_TEST) -v ./...

clean:
	$(GO_MOD) tidy
	$(GO_CLEAN)
	rm -rf $(OUT_DIR)/*

deps:
	$(GO_MOD) download

db_init_kootenay: build_kootenay
	$(OUT_DIR)/$(BINARY_NAME_KOOTENAY) dbinit
