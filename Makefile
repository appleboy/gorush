.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gorush-build"
# docker hub project name.
PRODUCTION_IMAGE := "gorush"
DEPLOY_ACCOUNT := "appleboy"
VERSION := $(shell git describe --tags || git rev-parse --short HEAD)
TARGETS_NOVENDOR := $(shell glide novendor)
export PROJECT_PATH = /go/src/github.com/appleboy/gorush

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
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

install:
	glide install

update:
	glide update

build_static:
	go build -ldflags="${EXTLDFLAGS}-s -w -X main.Version=${VERSION}" -o bin/gorush gorush.go

build: clean
	sh script/build.sh $(VERSION)

test: redis_test boltdb_test memory_test buntdb_test leveldb_test config_test
	go test -v -cover ./gorush/...

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

docker_build: clean
	tar -zcvf build.tar.gz gorush.go gorush config storage Makefile glide.lock glide.yaml
	sed -e "s/#VERSION#/$(VERSION)/g" docker/Dockerfile.build > docker/Dockerfile.tmp
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.tmp .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz

docker_production:
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

deploy: docker_build docker_production
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag $(PRODUCTION_IMAGE):latest $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)

docker_test: init clean
	docker-compose -p ${PRODUCTION_IMAGE} -f docker/docker-compose.testing.yml run gorush
	docker-compose -p ${PRODUCTION_IMAGE} -f docker/docker-compose.testing.yml down

fmt:
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

clean:
	-rm -rf build.tar.gz \
		gorush.tar.gz bin/* \
		gorush.tar.gz \
		gorush/gorush.db \
		storage/boltdb/gorush.db \
		.cover
