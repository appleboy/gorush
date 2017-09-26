FROM docker.int.eencloud.com/libeen:latest
ENV PKG_CONFIG_PATH /opt/een/lib/pkgconfig

ENV GOPATH /usr/src/app/go
ENV LD_LIBRARY_PATH /opt/een/lib

EXPOSE 8808

RUN apk update && \
    apk add go=1.8.3-r0 make git linux-headers musl-dev build-base gcc abuild binutils libc-dev

COPY ./ /usr/src/app/go/src/github.com/eencloud/gorush/
WORKDIR /usr/src/app/go/src/github.com/eencloud/gorush/
RUN mkdir /usr/src/app/go/src/github.com/eencloud/gorush/bin/
RUN mv /usr/src/app/go/src/github.com/eencloud/gorush/goeen /usr/src/app/go/src/github.com/eencloud/

RUN go get

RUN make docker_build

EXPOSE 8088
RUN chmod +x start.sh

CMD ["sh", "./start.sh"]
