.PHONY: all

VERSION=0.0.1

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gopush-build"

all: build

build: clean
	sh script/build.sh

docker_build: clean
	tar -zcvf build.tar.gz gopush.go gopush script
	docker build -t $(BUILD_IMAGE) -f docker/Dockerfile.build .
	docker run $(BUILD_IMAGE) > bin.tar.gz
	tar -zxvf bin.tar.gz
	-rm -rf bin.tar.gz build.tar.gz

test:
	cd gopush && go test -v -covermode=count -coverprofile=coverage.out

clean:
	rm -rf build.tar.gz bin.tar.gz bin/*
