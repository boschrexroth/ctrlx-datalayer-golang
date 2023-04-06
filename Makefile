PROJECT_NAME := "ctrlx-datalayer-golang"
PKG := "github.com/boschrexroth/$(PROJECT_NAME)"
PKG_LIST := ./pkg/datalayer ./pkg/sseclient ./pkg/token
COV_PKG_LIST := ./pkg/datalayer,./pkg/sseclient,./pkg/token
TST_LIST := ./test/datalayer ./test/sseclient ./test/token
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
BUILD := build/public/$(GOOS)_$(GOARCH)

DATALAYER_DEB_VERSION      := 1.20.0
DATALAYER_DEB_FILE_VERSION := 1.10.7

.PHONY: all go-dep apt-dep lint vet test test-coverage build clean

all: build

go-dep: ## Get go dependencies
	@go mod download
	@go mod vendor

apt-dep: ## Get native dependencies
	@sudo apt-get update
	@sudo apt-get install -y libsystemd-dev libzmq3-dev
	@wget --quiet https://github.com/boschrexroth/ctrlx-automation-sdk/releases/download/$(DATALAYER_DEB_VERSION)/ctrlx-datalayer-${DATALAYER_DEB_FILE_VERSION}.deb
	@sudo apt-get install -y -f ./ctrlx-datalayer-${DATALAYER_DEB_FILE_VERSION}.deb
	@rm ctrlx-datalayer-${DATALAYER_DEB_FILE_VERSION}.deb

lint: ## Lint Golang files
	@golangci-lint run

vet: ## Run go vet
	@go vet $(PKG_LIST)

test: ## Run unittests
	@go test -race -short -count=1 -mod=vendor $(TST_LIST)

memprofile: ## Run memory profiler
	@go test -race -short -count=1 -mod=vendor $(TST_LIST) -memprofile mem.pprof -memprofilerate 1
	@go tool pprof -http=:8080 mem.pprof 

cpuprofile: ## Run cpu profiler
	@go test -race -short -count=1 -mod=vendor $(TST_LIST) -cpuprofile cpu.pprof 
	@go tool pprof -http=:8080 cpu.pprof 

testcover: ## Run unittests with coverage 'go tool cover -html=coverage.out -o coverage.html'
	@go test -race -short -count=1 -mod=vendor -coverpkg=$(COV_PKG_LIST) -coverprofile=coverage.out -covermode=atomic $(TST_LIST)
	
build: go-dep ## Build the samples
	@CGO_ENABLED=1 go build -mod=vendor -o $(BUILD)/datalayer ./cmd/datalayer
	@CGO_ENABLED=1 go build -mod=vendor -o $(BUILD)/benchmark ./cmd/benchmark

clean: ## Remove previously build samples
	@rm -rf $(BUILD)/*

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
