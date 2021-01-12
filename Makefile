BIN ?= puller
ENTRY_POINT ?= ./cmd/puller/main.go
PKGS ?= cmd internal
EXPANDED_PKGS ?= ./cmd/... ./internal/...
# https://github.com/containers/image#building
BUILD_TAGS ?= containers_image_openpgp

run:
	@go run --tags $(BUILD_TAGS) $(ENTRY_POINT)

test:
	@go test $(EXPANDED_PKGS)

build:
	@go build --tags $(BUILD_TAGS) -o $(BIN) $(ENTRY_POINT)

format:
	@gofmt -l -s -w $(PKGS)

lint:
	@golangci-lint run

generate:
	@go generate $(EXPANDED_PKGS)

.PHONY: run test build format lint generate