.PHONY: all

VERSION=0.0.1

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gopush-build"

all: build

build: clean
	for GOOS in darwin linux windows; do \
		for GOARCH in 386 amd64; do \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -ldflags="-w" -o bin/$$GOOS/$$GOARCH/gopush gopush.go; \
		done \
	done

docker_build: clean
	tar -zcvf build.tar.gz gopush.go gopush
	docker build -t $(BUILD_IMAGE) -f docker/Dockerfile.build .
	docker run $(BUILD_IMAGE) > bin.tar.gz
	tar -zxvf bin.tar.gz
	-rm -rf bin.tar.gz build.tar.gz

clean:
	rm -rf build.tar.gz bin.tar.gz bin/*
