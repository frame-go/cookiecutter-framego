# Go version to require, run go version to see what version you are using
GO_VERSION := "go{{ cookiecutter.go_version }}"
GO_VERSION ?= $(GO_VERSION)
BUF_VENDOR_DIR := "./buf_vendor"

GO_BIN_DIR?=$(shell dirname `which go`)

.PHONY: all
all: fmt generate build

.PHONY: build
build: export GOBIN=$(CURDIR)/bin
build:
	go install -v $(CURDIR)/cmd/...

.PHONY: clean
clean:
	go clean
	rm -rf bin/*
	rm -rf tmp

.PHONY: generate
generate: buf-generate go-generate

.PHONY: go-generate
go-generate: export PATH:=$(CURDIR)/tools:$(PATH)
go-generate:
	go generate ./...

.PHONY: buf-generate
buf-generate: export PATH:=$(CURDIR)/tools:$(PATH)
buf-generate:
	@rm -rf $(BUF_VENDOR_DIR)
	buf generate
	buf export --output $(BUF_VENDOR_DIR)

.PHONY: lint
lint: buf-lint fmt go-lint

.PHONY: go-lint
go-lint: export PATH:=$(CURDIR)/tools:$(PATH)
go-lint:
	golangci-lint --version
	golangci-lint run

.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...

.PHONY: buf-lint
buf-lint:
	buf lint

.PHONY: check-go-version
check-go-version:
	@if ! go version | grep "$(GO_VERSION)" >/dev/null; then \
        printf "Wrong go version: "; \
        go version; \
        echo "Requires go version: $(GO_VERSION)"; \
        exit 2; \
    fi

.PHONY: init
init: export GOBIN=$(CURDIR)/tools
init: export PATH:=$(CURDIR)/tools:$(PATH)
init: check-go-version
	go mod tidy
	go install github.com/envoyproxy/protoc-gen-validate
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/swaggo/swag/cmd/swag
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/frame-go/protoc-gen-framego
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(CURDIR)/tools
	buf mod update
	buf generate
	buf export --output $(BUF_VENDOR_DIR)
	go mod tidy
	go fmt ./...
