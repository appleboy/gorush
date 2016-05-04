.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gorush-build"
TEST_IMAGE := "gorush-testing"
PRODUCTION_IMAGE := "gorush"
DEPLOY_ACCOUNT := "appleboy"
VERSION := $(shell git describe --tags)

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
	tar -zcvf build.tar.gz gorush.go gorush
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.build .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

docker_test:
	-docker rm -f gorush-redis
	@docker build --rm -t $(TEST_IMAGE) -f docker/Dockerfile.testing .
	@docker run --name gorush-redis -d redis
	@docker run --rm --link gorush-redis:redis -e ANDROID_TEST_TOKEN=$(ANDROID_TEST_TOKEN) -e ANDROID_API_KEY=$(ANDROID_API_KEY) $(TEST_IMAGE) sh -c "make test"
	@docker rm -f gorush-redis

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
