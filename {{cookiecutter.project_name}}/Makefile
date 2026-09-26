GO_VERSION := "go{{ cookiecutter.go_version }}"
GOLANGCI_LINT_VERSION := v2.14.0
BUF_VENDOR_DIR := "./buf_vendor"
GOLANGCI_LINT := $(CURDIR)/tools/golangci-lint

.PHONY: all
all: generate fmt build

.PHONY: build
build: export GOBIN=$(CURDIR)/bin
build:
	go install -v $(CURDIR)/cmd/...

.PHONY: run
run: build
	bin/{{ cookiecutter.app_package_name }} -c configs/debug/config.yaml

.PHONY: dev
dev:
	air

.PHONY: clean
clean:
	go clean
	rm -rf bin/*
	rm -rf tmp

.PHONY: generate
generate: buf-generate go-generate

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: buf-generate
buf-generate:
	@rm -rf $(BUF_VENDOR_DIR)
	buf generate
	buf export --output $(BUF_VENDOR_DIR)

.PHONY: lint
lint: buf-lint fmt go-lint

.PHONY: go-lint
go-lint:
	$(GOLANGCI_LINT) --version
	$(GOLANGCI_LINT) run --allow-serial-runners

.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...

.PHONY: test
test:
	go test -count=1 ./...

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
init: check-go-version
	go mod tidy
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/$(GOLANGCI_LINT_VERSION)/install.sh | sh -s -- -b $(CURDIR)/tools $(GOLANGCI_LINT_VERSION)
	buf dep update
	buf generate
	buf export --output $(BUF_VENDOR_DIR)
	go mod tidy
	go fmt ./...
