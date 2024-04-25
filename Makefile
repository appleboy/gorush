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

init:
#ifeq ($(ANDROID_TEST_TOKEN),)
#	@echo "Missing ANDROID_TEST_TOKEN Parameter"
#	@exit 1
#endif
#	@echo "Already set ANDROID_API_KEY and ANDROID_TEST_TOKEN globale variable."

vet:
	$(GO) vet ./...

embedmd:
	@hash embedmd > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/campoy/embedmd@master; \
	fi
	embedmd -d *.md

.PHONY: install
install: $(GOFILES)
	$(GO) install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'
	@echo "\n==>\033[32m Installed gorush to ${GOPATH}/bin/gorush\033[m"

.PHONY: build
build: $(EXECUTABLE)

.PHONY: $(EXECUTABLE)
$(EXECUTABLE): $(GOFILES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/$@

.PHONY: test
test: init
	@$(GO) test -v -cover -tags $(TAGS) -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

release: release-dirs release-build release-copy release-compress release-check

release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

release-build:
	@hash gox > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/mitchellh/gox@v1.0.1; \
	fi
	gox -os="$(TARGETS)" -arch="$(ARCHS)" -tags="$(TAGS)" -ldflags="$(EXTLDFLAGS)-s -w $(LDFLAGS)" -output="$(DIST)/binaries/$(EXECUTABLE)-$(VERSION)-{{.OS}}-{{.Arch}}"

.PHONY: release-compress
release-compress:
	@hash gxz > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/ulikunitz/xz/cmd/gxz@v0.5.10; \
	fi
	cd $(DIST)/release/; for file in `find . -type f -name "*"`; do echo "compressing $${file}" && gxz -k -9 $${file}; done;

release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/amd64/$(DEPLOY_IMAGE)

build_linux_i386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/i386/$(DEPLOY_IMAGE)

build_linux_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm64/$(DEPLOY_IMAGE)

build_linux_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm/$(DEPLOY_IMAGE)

build_linux_lambda:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags 'lambda' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/lambda/$(DEPLOY_IMAGE)

build_darwin_amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/amd64/$(DEPLOY_IMAGE)

build_darwin_i386:
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/i386/$(DEPLOY_IMAGE)

build_darwin_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/arm64/$(DEPLOY_IMAGE)

build_darwin_arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm GOARM=7 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/arm/$(DEPLOY_IMAGE)

build_darwin_lambda:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags 'lambda' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/darwin/lambda/$(DEPLOY_IMAGE)

clean:
	$(GO) clean -modcache -x -i ./...
	find . -name coverage.txt -delete
	find . -name *.tar.gz -delete
	find . -name *.db -delete
	-rm -rf release dist .cover

.PHONY: proto_install
proto_install:
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO)
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC)

generate_proto_js:
	npm install grpc-tools
	protoc -I rpc/proto rpc/proto/gorush.proto --js_out=import_style=commonjs,binary:rpc/example/node/ --grpc_out=rpc/example/node/ --plugin=protoc-gen-grpc="node_modules/.bin/grpc_tools_node_protoc_plugin"

generate_proto_go:
	protoc -I rpc/proto rpc/proto/gorush.proto --go_out=rpc/proto --go-grpc_out=require_unimplemented_servers=false:rpc/proto

generate_proto: generate_proto_go generate_proto_js

# install air command
.PHONY: air
air:
	@hash air > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/cosmtrek/air@latest; \
	fi

# run air
.PHONY: dev
dev: air
	air --build.cmd "make" --build.bin release/gorush

version:
	@echo $(VERSION)
