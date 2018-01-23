DIST := dist
EXECUTABLE := gorush

GO ?= go
DEPLOY_ACCOUNT := appleboy
DEPLOY_IMAGE := $(EXECUTABLE)
GOFMT ?= gofmt "-s"

TARGETS ?= linux darwin windows
ARCHS ?= amd64 386
PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
SOURCES ?= $(shell find . -name "*.go" -type f)
TAGS ?=
LDFLAGS ?= -X 'main.Version=$(VERSION)'
TMPDIR := $(shell mktemp -d 2>/dev/null || mktemp -d -t 'tempdir')
NODE_PROTOC_PLUGIN := $(shell which grpc_tools_node_protoc_plugin)
GOVENDOR := $(GOPATH)/bin/govendor
GOX := $(GOPATH)/bin/gox
MISSPELL := $(GOPATH)/bin/misspell

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

.PHONY: all
all: build

init:
ifeq ($(ANDROID_API_KEY),)
	@echo "Missing ANDROID_API_KEY Parameter"
	@exit 1
endif
ifeq ($(ANDROID_TEST_TOKEN),)
	@echo "Missing ANDROID_TEST_TOKEN Parameter"
	@exit 1
endif
	@echo "Already set ANDROID_API_KEY and ANDROID_TEST_TOKEN globale variable."

$(GOVENDOR):
	$(GO) get -u github.com/kardianos/govendor

$(GOX):
	$(GO) get -u github.com/mitchellh/gox

$(MISSPELL):
	$(GO) get -u github.com/client9/misspell/cmd/misspell

fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

vet:
	$(GO) vet $(PACKAGES)

deps:
	$(GO) get github.com/campoy/embedmd

embedmd:
	embedmd -d *.md

errcheck:
	@hash errcheck > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

lint:
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

unconvert:
	@hash unconvert > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mdempsky/unconvert; \
	fi
	for PKG in $(PACKAGES); do unconvert -v $$PKG || exit 1; done;

.PHONY: install
install: $(SOURCES)
	$(GO) install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'
	@echo "==> Installed gorush ${GOPATH}/bin/gorush"

.PHONY: build
build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/$@

.PHONY: misspell-check
misspell-check: $(MISSPELL)
	$(MISSPELL) -error $(GOFILES)

.PHONY: misspell
misspell: $(MISSPELL)
	$(MISSPELL) -w $(GOFILES)

test: fmt-check
	for PKG in $(PACKAGES); do $(GO) test -v -cover -coverprofile $$GOPATH/src/$$PKG/coverage.txt $$PKG || exit 1; done;

.PHONY: test-vendor
test-vendor: $(GOVENDOR)
	$(GOVENDOR) list +unused | tee "$(TMPDIR)/wc-gitea-unused"
	[ $$(cat "$(TMPDIR)/wc-gitea-unused" | wc -l) -eq 0 ] || echo "Warning: /!\\ Some vendor are not used /!\\"

	$(GOVENDOR) list +outside | tee "$(TMPDIR)/wc-gitea-outside"
	[ $$(cat "$(TMPDIR)/wc-gitea-outside" | wc -l) -eq 0 ] || exit 1

	$(GOVENDOR) status || exit 1

redis_test: init
	$(GO) test -v -cover ./storage/redis/...

boltdb_test: init
	$(GO) test -v -cover ./storage/boltdb/...

memory_test: init
	$(GO) test -v -cover ./storage/memory/...

buntdb_test: init
	$(GO) test -v -cover ./storage/buntdb/...

leveldb_test: init
	$(GO) test -v -cover ./storage/leveldb/...

config_test: init
	$(GO) test -v -cover ./config/...

html:
	$(GO) tool cover -html=.cover/coverage.txt

release: release-dirs release-build release-copy release-check

release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

release-build: $(GOX)
	$(GOX) -os="$(TARGETS)" -arch="$(ARCHS)" -tags="$(TAGS)" -ldflags="$(EXTLDFLAGS)-s -w $(LDFLAGS)" -output="$(DIST)/binaries/$(EXECUTABLE)-$(VERSION)-{{.OS}}-{{.Arch}}"

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

docker_image:
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
	$(GO) clean -x -i ./...
	find . -name coverage.txt -delete
	find . -name *.tar.gz -delete
	find . -name *.db -delete
	-rm -rf release dist .cover

rpc/example/node/gorush_*_pb.js: rpc/proto/gorush.proto
	@hash grpc_tools_node_protoc_plugin > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		npm install -g grpc-tools; \
	fi
	protoc -I rpc/proto rpc/proto/gorush.proto --js_out=import_style=commonjs,binary:rpc/example/node/ --grpc_out=rpc/example/node/ --plugin=protoc-gen-grpc=$(NODE_PROTOC_PLUGIN)

rpc/proto/gorush.pb.go: rpc/proto/gorush.proto
	protoc -I rpc/proto rpc/proto/gorush.proto --go_out=plugins=grpc:rpc/proto

generate_proto: rpc/proto/gorush.pb.go rpc/example/node/gorush_*_pb.js

version:
	@echo $(VERSION)
