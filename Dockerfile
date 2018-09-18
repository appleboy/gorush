FROM docker.eencloud.com/goeen:latest

RUN apk update && apk add git go=1.9.4-r0 glib-dev libev-dev

WORKDIR /usr/src/app/go/src/github.com/eencloud/gorush/
COPY . .

RUN go get
RUN go build -o bin/gorush

CMD ["./start.sh"]
