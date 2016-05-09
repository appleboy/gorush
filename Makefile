.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gorush-build"
PRODUCTION_IMAGE := "gorush"
DEPLOY_ACCOUNT := "appleboy"
VERSION := $(shell git describe --tags)
TIMENAME := $(shell date '+%Y%m%d%H%M%S%s')

all: build

build: clean
	sh script/build.sh $(VERSION)

test: redis_test boltdb_test memory_test config_test
	go test -v -cover -covermode=count -coverprofile=coverage.out ./gorush/...

redis_test:
	go test -v -cover -covermode=count -coverprofile=coverage.out ./storage/redis/...

boltdb_test:
	go test -v -cover -covermode=count -coverprofile=coverage.out ./storage/boltdb/...

memory_test:
	go test -v -cover -covermode=count -coverprofile=coverage.out ./storage/memory/...

config_test:
	go test -v -cover -covermode=count -coverprofile=coverage.out ./config/...

html:
	go tool cover -html=coverage.out

docker_build: clean
	tar -zcvf build.tar.gz gorush.go gorush config storage Makefile glide.lock glide.yaml
	sed -e "s/#VERSION#/$(VERSION)/g" docker/Dockerfile.build > docker/Dockerfile.tmp
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.tmp .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

docker_test:
	@docker build --rm -t $(TIMENAME)-test -f docker/Dockerfile.testing .
	@docker run --name $(TIMENAME)-redis -d redis
	@docker run --rm --link $(TIMENAME)-redis:redis -e ANDROID_TEST_TOKEN=$(ANDROID_TEST_TOKEN) -e ANDROID_API_KEY=$(ANDROID_API_KEY) $(TIMENAME)-test sh -c "make test"
	@docker rm -f $(TIMENAME)-redis
	@docker rmi -f $(TIMENAME)-test

deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag -f $(PRODUCTION_IMAGE):latest $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)

bundle:
	glide install

bundle_update:
	glide update --all-dependencies --resolve-current

lint:
	golint gorush

clean:
	-rm -rf build.tar.gz gorush.tar.gz bin/*
