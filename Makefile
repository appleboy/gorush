.PHONY: all

DEPS := $(wildcard *.go)
BUILD_IMAGE := "gorush-build"
# docker hub project name.
PRODUCTION_IMAGE := "gorush"
DEPLOY_ACCOUNT := "appleboy"
VERSION := $(shell git describe --tags)
TARGETS_NOVENDOR := $(shell glide novendor)

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

coverage:
	sh ./script/coverage.sh testing atomic

test: redis_test boltdb_test memory_test config_test
	go test -v -cover ./gorush/...

redis_test: init
	go test -v -cover ./storage/redis/...

boltdb_test: init
	go test -v -cover ./storage/boltdb/...

memory_test: init
	go test -v -cover ./storage/memory/...

config_test: init
	go test -v -cover ./config/...

html:
	go tool cover -html=.cover/coverage.txt

docker_build: clean
	tar -zcvf build.tar.gz gorush.go gorush config storage Makefile glide.lock glide.yaml
	sed -e "s/#VERSION#/$(VERSION)/g" docker/Dockerfile.build > docker/Dockerfile.tmp
	docker build --rm -t $(BUILD_IMAGE) -f docker/Dockerfile.tmp .
	docker run --rm $(BUILD_IMAGE) > gorush.tar.gz

docker_test: init
	docker-compose -p ${PRODUCTION_IMAGE} -f docker/docker-compose.testing.yml run --rm gorush
	docker-compose -p ${PRODUCTION_IMAGE} -f docker/docker-compose.testing.yml down

docker_production: docker_build
	docker build --rm -t $(PRODUCTION_IMAGE) -f docker/Dockerfile.dist .

deploy: docker_production
ifeq ($(tag),)
	@echo "Usage: make $@ tag=<tag>"
	@exit 1
endif
	docker tag $(PRODUCTION_IMAGE):latest $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)
	docker push $(DEPLOY_ACCOUNT)/$(PRODUCTION_IMAGE):$(tag)

install:
	@which glide || (curl https://glide.sh/get | sh)
	@which go-junit-report || go get -u github.com/jstemmer/go-junit-report
	@which gocov || go get -u github.com/axw/gocov/gocov
	@which gocov-xml || go get -u github.com/AlekSi/gocov-xml
	@which golint || go get -u github.com/golang/lint/golint
	@glide install

update:
	@glide up

fmt:
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

lint:
	@golint -set_exit_status=1 ./...

vet:
	@go vet -n -x ./...

junit_report:
	sh ./script/coverage.sh junit

coverage_report:
	sh ./script/coverage.sh coverage

lint_report:
	sh ./script/coverage.sh lint

vet_report:
	sh ./script/coverage.sh vet

cloc_report:
	sh ./script/coverage.sh cloc

report: junit_report coverage_report lint_report vet_report cloc_report

clean:
	-rm -rf build.tar.gz \
		gorush.tar.gz bin/* \
		gorush.tar.gz \
		gorush/gorush.db \
		storage/boltdb/gorush.db \
		.cover
