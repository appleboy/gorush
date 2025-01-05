DIST := dist
EXECUTABLE := gorush

GO ?= go
DEPLOY_ACCOUNT := appleboy
DEPLOY_IMAGE := $(EXECUTABLE)
GOFMT ?= gofumpt -l -s -extra

TARGETS ?= linux darwin windows
ARCHS ?= amd64 386
PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
SOURCES ?= $(shell find . -name "*.go" -type f)
TAGS ?=
TMPDIR := $(shell mktemp -d 2>/dev/null || mktemp -d -t 'tempdir')
NODE_PROTOC_PLUGIN := $(shell which grpc_tools_node_protoc_plugin)

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
ifeq ($(FCM_CREDENTIAL),)
	@echo "Missing FCM_CREDENTIAL Parameter"
	@exit 1
endif
ifeq ($(FCM_TEST_TOKEN),)
	@echo "Missing FCM_TEST_TOKEN Parameter"
	@exit 1
endif
	@echo "Already set FCM_CREDENTIAL and endif global variable."

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

docker_build:
	GOOS=linux GOARCH=amd64 $(GO) build -a -tags '$(TAGS)'  -o bin/$(EXECUTABLE)

docker_build_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO) build -a -tags '$(TAGS)' -ldflags "$(EXTLDFLAGS)-s -w $(LDFLAGS)" -o bin/$(EXECUTABLE)-arm64

docker_build_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 $(GO) build -a -tags '$(TAGS)' -ldflags "$(EXTLDFLAGS)-s -w $(LDFLAGS)" -o bin/$(EXECUTABLE)-arm

docker_image:
	rm -rf goeen
	git clone git@github.com:EENCloud/goeen.git goeen
	docker build -t $(DEPLOY_ACCOUNT)/$(DEPLOY_IMAGE) -f Dockerfile .

docker_release: docker_image

docker_deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag $(DEPLOY_ACCOUNT)/$(EXECUTABLE):latest $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag)

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
