.PHONY: all test build fmt vet errcheck lint install update release-dirs release-build release-copy release-check release

export PROJECT_PATH = /go/src/github.com/appleboy/gorush

DIST := dist
EXECUTABLE := gorush

BUILD_IMAGE := "gorush-build"
DEPLOY_ACCOUNT := appleboy
DEPLOY_IMAGE := $(EXECUTABLE)

TARGETS ?= linux darwin
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
SOURCES ?= $(shell find . -name "*.go" -type f)
TAGS ?=
LDFLAGS += -X 'main.Version=$(VERSION)'

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(DRONE_BRANCH)
	else
		VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
	endif
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
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

errcheck:
	@which errcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

dep_install:
	glide install

dep_update:
	glide up

install: $(wildcard *.go)
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
	@which gox > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/mitchellh/gox; \
	fi
	gox -os="$(TARGETS)" -arch="amd64 386" -tags="$(TAGS)" -ldflags="-s -w $(LDFLAGS)" -output="$(DIST)/binaries/$(EXECUTABLE)-$(VERSION)-{{.OS}}-{{.Arch}}"

release-copy:
	$(foreach file,$(wildcard $(DIST)/binaries/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)


docker_binary_build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o bin/$@

docker_image:
	docker build -t $(DEPLOY_ACCOUNT)/$(DEPLOY_IMAGE) -f Dockerfile .

docker_release: docker_binary_build docker_image

docker_build:
	tar -zcvf build.tar.gz gorush.go gorush config storage Makefile glide.lock glide.yaml
	sed -e "s/#VERSION#/$(VERSION)/g" docker/Dockerfile.build > docker/Dockerfile.tmp
	docker build -t $(BUILD_IMAGE) -f docker/Dockerfile.tmp .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz

docker_production:
	docker build -t $(EXECUTABLE) -f docker/Dockerfile.dist .

docker_deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag $(EXECUTABLE):latest $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(EXECUTABLE):$(tag)

docker_test: init clean
	docker-compose -p ${EXECUTABLE} -f docker/docker-compose.testing.yml run gorush
	docker-compose -p ${EXECUTABLE} -f docker/docker-compose.testing.yml down

clean:
	go clean -x -i ./...
	find . -name coverage.txt -delete
	-rm -rf build.tar.gz \
		gorush.tar.gz bin/* \
		gorush.tar.gz \
		gorush/gorush.db \
		storage/boltdb/gorush.db \
		.cover

version:
	@echo $(VERSION)
