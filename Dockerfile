FROM docker.int.eencloud.com/goeen:52

WORKDIR /usr/src/go/src/github.com/eencloud/gorush/
COPY . .

RUN go get
RUN go build -o bin/gorush

CMD ["./start.sh"]
