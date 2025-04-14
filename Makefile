EXECUTABLE := gorush
GO ?= go
GOFMT ?= gofumpt -l -s -extra
GOFILES := $(shell find . -name "*.go" -type f)
TAGS ?= sqlite
LDFLAGS ?= -X main.version=$(VERSION) -X main.commit=$(COMMIT)

PROTOC_GEN_GO=v1.36.6
PROTOC_GEN_GO_GRPC=v1.5.1

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
endif

COMMIT ?= $(shell git rev-parse --short HEAD)

all: build

.PHONY: help
help: ## Print this help message.
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: ## check the FCM_CREDENTIAL and FCM_TEST_TOKEN
	@echo "==> Check FCM_CREDENTIAL and FCM_TEST_TOKEN"
ifeq ($(FCM_CREDENTIAL),)
	@echo "Missing FCM_CREDENTIAL Parameter"
	@exit 1
endif
ifeq ($(FCM_TEST_TOKEN),)
	@echo "Missing FCM_TEST_TOKEN Parameter"
	@exit 1
endif
	@echo "Already set FCM_CREDENTIAL and endif global variable."

.PHONY: install ## Install the gorush binary
install: $(GOFILES)
	$(GO) install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'
	@echo "\n==>\033[32m Installed gorush to ${GOPATH}/bin/gorush\033[m"

.PHONY: build ## Build the gorush binary
build: $(EXECUTABLE)

.PHONY: $(EXECUTABLE)
$(EXECUTABLE): $(GOFILES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/$@

.PHONY: test ## Run the tests
test: init
	@$(GO) test -v -cover -tags $(TAGS) -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

build_linux_amd64: ## build the gorush binary for linux amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/amd64/$(EXECUTABLE)

build_linux_i386: ## build the gorush binary for linux i386
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/i386/$(EXECUTABLE)

build_linux_arm64: ## build the gorush binary for linux arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm64/$(EXECUTABLE)

build_linux_arm: ## build the gorush binary for linux arm
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm/$(EXECUTABLE)

build_linux_lambda: ## build the gorush binary for linux lambda
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags 'lambda' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/lambda/$(EXECUTABLE)

build_darwin_amd64: ## build the gorush binary for darwin amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/amd64/$(EXECUTABLE)

build_darwin_i386: ## build the gorush binary for darwin i386
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/i386/$(EXECUTABLE)

build_darwin_arm64: ## build the gorush binary for darwin arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/arm64/$(EXECUTABLE)

build_darwin_arm: ## build the gorush binary for darwin arm
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm GOARM=7 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/arm/$(EXECUTABLE)

build_darwin_lambda: ## build the gorush binary for darwin lambda
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags 'lambda' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/lambda/$(EXECUTABLE)

clean: ## Clean the build
	$(GO) clean -modcache -x -i ./...
	find . -name coverage.txt -delete
	find . -name *.tar.gz -delete
	find . -name *.db -delete
	-rm -rf release dist .cover

.PHONY: proto_install
proto_install: ## install the protoc-gen-go and protoc-gen-go-grpc
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO)
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC)

generate_proto_js: ## generate the proto file for nodejs
	npm install grpc-tools
	protoc -I rpc/proto rpc/proto/gorush.proto --js_out=import_style=commonjs,binary:rpc/example/node/ --grpc_out=rpc/example/node/ --plugin=protoc-gen-grpc="node_modules/.bin/grpc_tools_node_protoc_plugin"

generate_proto_go: ## generate the proto file for golang
	protoc -I rpc/proto rpc/proto/gorush.proto --go_out=rpc/proto --go-grpc_out=require_unimplemented_servers=false:rpc/proto

generate_proto: generate_proto_go generate_proto_js

.PHONY: air
air: ## install air for hot reload
	@hash air > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/cosmtrek/air@latest; \
	fi

.PHONY: dev ## run the air for hot reload
dev: air
	air --build.cmd "make" --build.bin release/gorush

version: ## print the version
	@echo $(VERSION)
