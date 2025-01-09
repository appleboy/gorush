DIST := dist
EXECUTABLE := gorush

GO ?= go
DEPLOY_ACCOUNT := appleboy
DEPLOY_IMAGE := $(EXECUTABLE)
GOFMT ?= gofumpt -l -s -extra

TARGETS ?= linux darwin windows
ARCHS ?= amd64
GOFILES := $(shell find . -name "*.go" -type f)
TAGS ?= sqlite
LDFLAGS ?= -X main.version=$(VERSION) -X main.commit=$(COMMIT)

PROTOC_GEN_GO=v1.28
PROTOC_GEN_GO_GRPC=v1.2

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

.PHONY: all
all: build

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## init: check the environment variables
init:
ifeq ($(FCM_CREDENTIAL),)
	@echo "Missing FCM_CREDENTIAL Parameter"
	@exit 1
endif
ifeq ($(FCM_TEST_TOKEN),)
	@echo "Missing FCM_TEST_TOKEN Parameter"
	@exit 1
endif
	@echo "Already set FCM_CREDENTIAL and endif global variable."

## install: install the gorush binary
.PHONY: install
install: $(GOFILES)
	$(GO) install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'
	@echo "\n==>\033[32m Installed gorush to ${GOPATH}/bin/gorush\033[m"

## build: build the gorush binary
.PHONY: build
build: $(EXECUTABLE)

.PHONY: $(EXECUTABLE)
$(EXECUTABLE): $(GOFILES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/$@

## test: run the tests
.PHONY: test
test: init
	@$(GO) test -v -cover -tags $(TAGS) -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

## build_linux_amd64: build the gorush binary for linux amd64
build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/amd64/$(DEPLOY_IMAGE)

## build_linux_i386: build the gorush binary for linux i386
build_linux_i386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/i386/$(DEPLOY_IMAGE)

## build_linux_arm64: build the gorush binary for linux arm64
build_linux_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm64/$(DEPLOY_IMAGE)

## build_linux_arm: build the gorush binary for linux arm
build_linux_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm/$(DEPLOY_IMAGE)

## build_linux_lambda: build the gorush binary for linux lambda
build_linux_lambda:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags 'lambda' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/lambda/$(DEPLOY_IMAGE)

## build_darwin_amd64: build the gorush binary for darwin amd64
build_darwin_amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/amd64/$(DEPLOY_IMAGE)

## build_darwin_i386: build the gorush binary for darwin i386
build_darwin_i386:
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/i386/$(DEPLOY_IMAGE)

## build_darwin_arm64: build the gorush binary for darwin arm64
build_darwin_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/arm64/$(DEPLOY_IMAGE)

## build_darwin_arm: build the gorush binary for darwin arm
build_darwin_arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm GOARM=7 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/arm/$(DEPLOY_IMAGE)

## build_darwin_lambda: build the gorush binary for darwin lambda
build_darwin_lambda:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags 'lambda' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/lambda/$(DEPLOY_IMAGE)

## clean: cleans up the project directory
#   Cleans up the project directory by performing the following actions:
#   - Runs `go clean` with the `-modcache`, `-x`, and `-i` flags to clean the module cache and remove installed packages.
#   - Deletes all files named `coverage.txt` in the project directory and its subdirectories.
#   - Deletes all files with the `.tar.gz` extension in the project directory and its subdirectories.
#   - Deletes all files with the `.db` extension in the project directory and its subdirectories.
#   - Removes the `release`, `dist`, and `.cover` directories if they exist.
clean:
	$(GO) clean -modcache -x -i ./...
	find . -name coverage.txt -delete
	find . -name *.tar.gz -delete
	find . -name *.db -delete
	-rm -rf release dist .cover

## proto_install: install the protoc-gen-go and protoc-gen-go-grpc
.PHONY: proto_install
proto_install:
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO)
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC)

## generate_proto_js: generate the proto file for nodejs
generate_proto_js:
	npm install grpc-tools
	protoc -I rpc/proto rpc/proto/gorush.proto --js_out=import_style=commonjs,binary:rpc/example/node/ --grpc_out=rpc/example/node/ --plugin=protoc-gen-grpc="node_modules/.bin/grpc_tools_node_protoc_plugin"

## generate_proto_go: generate the proto file for golang
generate_proto_go:
	protoc -I rpc/proto rpc/proto/gorush.proto --go_out=rpc/proto --go-grpc_out=require_unimplemented_servers=false:rpc/proto

## generate_proto: generate the proto file for golang and nodejs
generate_proto: generate_proto_go generate_proto_js

## air: install air for hot reload
.PHONY: air
air:
	@hash air > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/cosmtrek/air@latest; \
	fi

## dev: run the air for hot reload
.PHONY: dev
dev: air
	air --build.cmd "make" --build.bin release/gorush

## version: print the version
version:
	@echo $(VERSION)
