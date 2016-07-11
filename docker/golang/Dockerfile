FROM golang:1.6.2-alpine

MAINTAINER Bo-Yi Wu <appleboy.tw@gmail.com>

RUN apk update \
  && apk add git make curl && rm -rf /var/cache/apk/*

# install glide package management.
RUN curl https://glide.sh/get | sh

RUN mkdir /app && mkdir -p $GOPATH/src/github.com/appleboy
RUN ln -sf /app $GOPATH/src/github.com/appleboy/gorush

WORKDIR $GOPATH/src/github.com/appleboy/gorush
