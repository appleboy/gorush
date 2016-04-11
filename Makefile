.PHONY: all

VERSION=0.0.1

DEPS := $(wildcard *.go)
BUILD_IMAGE := "appleboy/gopush-build:latest"
TEST_IMAGE := "appleboy/gopush-testing:latest"

all: build

build: clean
	sh script/build.sh

test:
	cd gopush && go test -v -covermode=count -coverprofile=coverage.out

docker_build: clean
	tar -zcvf build.tar.gz gopush.go gopush script
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.build .
	docker run --rm $(BUILD_IMAGE) > bin.tar.gz
	tar -zxvf bin.tar.gz
	-rm -rf bin.tar.gz build.tar.gz

docker_test:
	@docker build --rm -t $(TEST_IMAGE) -f docker/Dockerfile.testing .
	@docker run --rm -e ANDROID_TEST_TOKEN=$(ANDROID_TEST_TOKEN) -e ANDROID_API_KEY=$(ANDROID_API_KEY) $(TEST_IMAGE) sh -c "cd gopush && go test -v"

clean:
	rm -rf build.tar.gz bin.tar.gz bin/*
