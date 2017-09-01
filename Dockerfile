FROM centos:6
RUN rpm -Uvh http://download.fedoraproject.org/pub/epel/6/i386/epel-release-6-8.noarch.rpm
RUN yum -y install golang git
ENV GOPATH=/usr/src/app/go
ENV GOBIN=/usr/src/app/go/bin/
ENV GODEBUG=netdns=go


ADD config/config.yml /

COPY . /usr/src/app/go/gorush/
WORKDIR /usr/src/app/go/gorush/
RUN mkdir /usr/src/app/go/bin/

RUN go get && make docker_build

EXPOSE 8088
RUN chmod +x start.sh

CMD ["sh", "./start.sh"]
