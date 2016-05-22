.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gorush-build"
PRODUCTION_IMAGE := "gorush"
DEPLOY_ACCOUNT := "appleboy"
VERSION := $(shell git describe --tags)
RANDOM := $(shell date '+%Y%m%d%H%M%S%s')

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

build: clean
	sh script/build.sh $(VERSION)

test: redis_test boltdb_test memory_test config_test
	go test -v -cover -covermode=count -coverprofile=coverage.out ./gorush/...

redis_test: init
	go test -v -cover -covermode=count -coverprofile=coverage.out ./storage/redis/...

boltdb_test: init
	go test -v -cover -covermode=count -coverprofile=coverage.out ./storage/boltdb/...

memory_test: init
	go test -v -cover -covermode=count -coverprofile=coverage.out ./storage/memory/...

config_test: init
	go test -v -cover -covermode=count -coverprofile=coverage.out ./config/...

html:
	go tool cover -html=coverage.out && unlink coverage.out

docker_build: clean
	tar -zcvf build.tar.gz gorush.go gorush config storage Makefile glide.lock glide.yaml
	sed -e "s/#VERSION#/$(VERSION)/g" docker/Dockerfile.build > docker/Dockerfile.tmp
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.tmp .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

docker_test: init
	docker-compose -p ${RANDOM} -f docker/docker-compose.testing.yml run --rm gorush
	docker-compose -p ${RANDOM} -f docker/docker-compose.testing.yml down

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
	-rm -rf build.tar.gz gorush.tar.gz bin/* coverage.out gorush.tar.gz gorush/gorush.db storage/boltdb/gorush.db
