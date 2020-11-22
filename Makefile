.PHONY: all build build_columbia build_kootenay clean config_kootenay db_init_kootenay deps proto test
os = $(shell uname -s | tr 'A-Z' 'a-z')
march = $(shell uname -m)

BINARY_NAME_COLUMBIA ?= columbia
BINARY_NAME_KOOTENAY ?= kootenay
BUILD_ROOT_COLUMBIA ?= ./cmd/columbia
BUILD_ROOT_KOOTENAY ?= ./cmd/kootenay
GO_CMD ?= go
GOPATH = $(shell ${GO_CMD} env GOPATH)
export PATH := ${GOPATH}/bin:${PATH}
GO_BUILD = ${GO_CMD} build
GO_CLEAN = ${GO_CMD} clean
GO_TEST = ${GO_CMD} test
GO_GET = ${GO_CMD} get
GO_MOD = ${GO_CMD} mod
OUT_DIR ?= ./dist
PROTOC ?= third_party/bin/protoc-${os}-${march}
PROTO_PATH = api/proto/v0


all: proto test build

build: build_columbia build_kootenay

build_columbia:
	CGO_ENABLED=0 $(GO_BUILD) -o $(OUT_DIR)/$(BINARY_NAME_COLUMBIA) -v ./cmd/columbia

build_kootenay:
	CGO_ENABLED=0 $(GO_BUILD) -o $(OUT_DIR)/$(BINARY_NAME_KOOTENAY) -v ./cmd/kootenay

clean:
	$(GO_MOD) tidy
	$(GO_MOD) vendor
	$(GO_CLEAN)
	rm -rf $(OUT_DIR)/*
	rm -f config.toml

config_kootenay:
	cp -f configs/kootenay.config.toml config.toml

db_init_kootenay: build_kootenay
	$(OUT_DIR)/$(BINARY_NAME_KOOTENAY) dbinit

deps:
	$(GO_MOD) download
	$(GO_MOD) vendor

proto:
	$(GO_GET) google.golang.org/protobuf/cmd/protoc-gen-go@v1
	$(GO_GET) google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1
	$(GO_MOD) vendor
	$(PROTOC) --proto_path=$(PROTO_PATH) --proto_path=third_party/include --go_out=. $(PROTO_PATH)/*.proto
	$(PROTOC) --proto_path=$(PROTO_PATH) --proto_path=third_party/include --go-grpc_out=. $(PROTO_PATH)/*.proto

test: proto
	$(GO_TEST) -v ./...
