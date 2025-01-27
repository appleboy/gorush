FROM harbor.eencloud.com/vms/goeen:latest

# Install required tools and libraries
RUN apk update && apk add --no-cache git glib-dev libev-dev curl

# Download and install Go 1.23.4 manually
RUN curl -sSL https://dl.google.com/go/go1.23.4.linux-amd64.tar.gz | tar -C /usr/local -xz && \
    ln -s /usr/local/go/bin/* /usr/bin/

WORKDIR /usr/src/app/go/src/github.com/eencloud/gorush/
COPY . .

RUN go mod tidy
RUN go build -o bin/gorush

CMD ["./start.sh"]
