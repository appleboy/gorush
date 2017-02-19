.PHONY: all example test

TARGETS_NOVENDOR := $(shell glide novendor)
export PROJECT_PATH = /go/src/github.com/appleboy/gofight

all: install test

install:
	go get -t -d -v ./...

example:
	cd example && go test -v -cover .

test: example
	go test -v -cover .

docker_test: clean
	docker run --rm \
		-v $(PWD):$(PROJECT_PATH) \
		-w=$(PROJECT_PATH) \
		appleboy/golang-testing \
		sh -c "make install && coverage all"

clean:
	rm -rf .cover vendor
