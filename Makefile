.PHONY: all gorush test build fmt vet errcheck lint install update release-dirs release-build release-copy release-check release

DIST := dist
EXECUTABLE := gorush

DEPLOY_ACCOUNT := jaraxasoftware
DEPLOY_IMAGE := $(EXECUTABLE)
CONTAINER := js-gorush
PORT := 8088

TARGETS ?= linux darwin windows
ARCHS ?= amd64 386
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)
TAGS ?=
LDFLAGS ?= -X 'main.Version=$(VERSION)'

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

fmt:
	find . -name "*.go" -type f -not -path "./vendor/*" | xargs gofmt -s -w

vet:
	go vet $(PACKAGES)

deps:
	go get github.com/campoy/embedmd

embedmd:
	embedmd -d *.md

errcheck:
	@hash errcheck > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

lint:
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

unconvert:
	@hash unconvert > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/mdempsky/unconvert; \
	fi
	for PKG in $(PACKAGES); do unconvert -v $$PKG || exit 1; done;

install: $(SOURCES)
	go install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'

build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	go build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o bin/$@

test:
	for PKG in $(PACKAGES); do go test -v -cover -coverprofile $$GOPATH/src/$$PKG/coverage.txt $$PKG || exit 1; done;

redis_test: init
	go test -v -cover ./storage/redis/...

boltdb_test: init
	go test -v -cover ./storage/boltdb/...

memory_test: init
	go test -v -cover ./storage/memory/...

buntdb_test: init
	go test -v -cover ./storage/buntdb/...

leveldb_test: init
	go test -v -cover ./storage/leveldb/...

config_test: init
	go test -v -cover ./config/...

html:
	go tool cover -html=.cover/coverage.txt

release: release-dirs release-build release-copy release-check

release-dirs:
	mkdir -p $(DIST)/binaries $(DIST)/release

release-build:
	@hash gox > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/mitchellh/gox; \
	fi
	gox -os="$(TARGETS)" -arch="$(ARCHS)" -tags="$(TAGS)" -ldflags="$(EXTLDFLAGS)-s -w $(LDFLAGS)" -output="$(DIST)/binaries/$(EXECUTABLE)-$(VERSION)-{{.OS}}-{{.Arch}}"

release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

docker_build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags "$(EXTLDFLAGS)-s -w $(LDFLAGS)" -o bin/$(EXECUTABLE)

docker_image:
	docker build -t $(DEPLOY_ACCOUNT)/$(DEPLOY_IMAGE) -f Dockerfile .

docker_release: docker_build docker_image

docker_deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag $(DEPLOY_ACCOUNT)/$(EXECUTABLE):latest $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag)

docker_stop:
	@if [ $(shell docker ps -a | grep -ci $(CONTAINER)) -eq 1 ]; then \
		docker stop $(CONTAINER) > /dev/null 2>&1; \
	fi

docker_rm: docker_stop
	@if [ $(shell docker ps -a | grep -ci $(CONTAINER)) -eq 1 ]; then \
		docker rm $(CONTAINER) > /dev/null 2>&1; \
	fi

docker_rmi: docker_rm
ifeq ($(tag),)
	@if [ $(shell docker images | grep -ci $(DEPLOY_ACCOUNT)/$(EXECUTABLE)) -eq 1 ]; then \
		docker rmi $(DEPLOY_ACCOUNT)/$(EXECUTABLE):latest > /dev/null 2>&1; \
	fi;
	@if [ $(shell docker images | grep -ci centurylink/ca-certs) -eq 1 ]; then \
		docker rmi centurylink/ca-certs:latest > /dev/null 2>&1; \
	fi
else
	@if [ $(shell docker images | grep -ci $(DEPLOY_ACCOUNT)/$(EXECUTABLE)) -eq 1 ]; then \
		docker rmi $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag) > /dev/null 2>&1; \
	fi;
	@if [ $(shell docker images | grep -ci centurylink/ca-certs) -eq 1 ]; then \
		docker rmi centurylink/ca-certs:latest > /dev/null 2>&1; \
	fi
endif

docker_run: docker_rm
	docker run -ti -d --name $(CONTAINER) --restart always \
	-p ${PORT}:8088 \
	-v ${CURDIR}/config:/config:ro \
	$(DEPLOY_ACCOUNT)/$(EXECUTABLE):latest /gorush -c /config/config.yml

docker_test:
	curl \
	-XGET \
	-H "Accept: application/json" \
 	"localhost:$(PORT)/api/stat/go" | python -mjson.tool

docker_save:
	docker save $(DEPLOY_ACCOUNT)/$(EXECUTABLE) | gzip > $(DEPLOY_ACCOUNT)_$(EXECUTABLE).tar.gz

docker_load:
	gunzip < $(DEPLOY_ACCOUNT)_$(EXECUTABLE).tar.gz | docker load

coverage:
	curl -s https://codecov.io/bash > .codecov && \
	chmod +x .codecov && \
	./.codecov -f .cover/coverage.txt

clean:
	go clean -x -i ./...
	find . -name coverage.txt -delete
	find . -name *.tar.gz -delete
	find . -name *.db -delete
	-rm -rf bin/* \
		.cover

version:
	@echo $(VERSION)
