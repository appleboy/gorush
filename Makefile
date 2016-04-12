.PHONY: all docker_build deploy

VERSION=0.0.1

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gopush-build"
TEST_IMAGE := "gopush-testing"
PRODUCTION_IMAGE := "gopush"
DEPLOY_ACCOUNT := "appleboy"

all: build

build: clean
	sh script/build.sh

test:
	cd gopush && go test -v -covermode=count -coverprofile=coverage.out

docker_build: clean
	tar -zcvf build.tar.gz gopush.go gopush
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.build .
	docker run --rm $(BUILD_IMAGE) > gopush.tar.gz
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

docker_test:
	@docker build --rm -t $(TEST_IMAGE) -f docker/Dockerfile.testing .
	@docker run --rm -e ANDROID_TEST_TOKEN=$(ANDROID_TEST_TOKEN) -e ANDROID_API_KEY=$(ANDROID_API_KEY) $(TEST_IMAGE) sh -c "cd gopush && go test -v"

deploy:
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag -f $(PRODUCTION_IMAGE):latest $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)

clean:
	-rm -rf build.tar.gz gopush.tar.gz bin/*
