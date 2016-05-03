.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gorush-build"
TEST_IMAGE := "gorush-testing"
PRODUCTION_IMAGE := "gorush"
DEPLOY_ACCOUNT := "appleboy"

all: build

build: clean
	sh script/build.sh

test: memory_test redis_test boltdb_test
	cd gorush && go test -cover -v -coverprofile=coverage.out

memory_test:
	cd storage/memory && go test -v -cover *.go

redis_test:
	cd storage/redis && go test -v -cover *.go

boltdb_test:
	cd storage/boltdb && go test -v -cover *.go

html: test
	cd gorush && go tool cover -html=coverage.out

docker_build: clean
	tar -zcvf build.tar.gz gorush.go gorush
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.build .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

docker_test:
	@docker build --rm -t $(TEST_IMAGE) -f docker/Dockerfile.testing .
	@docker run --rm -e ANDROID_TEST_TOKEN=$(ANDROID_TEST_TOKEN) -e ANDROID_API_KEY=$(ANDROID_API_KEY) $(TEST_IMAGE) sh -c "cd gorush && go test -v"

deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag -f $(PRODUCTION_IMAGE):latest $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)

lint:
	golint gorush

clean:
	-rm -rf build.tar.gz gorush.tar.gz bin/*
