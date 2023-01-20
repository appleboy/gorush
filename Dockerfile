FROM golang:1.16 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags '-X main.version=v1.13.0 -X main.build=1' -a -o bin/gorush

FROM alpine:latest AS final
COPY --from=builder /app/bin/gorush /bin/gorush
EXPOSE 8088 9000
ENTRYPOINT ["/bin/gorush"]
