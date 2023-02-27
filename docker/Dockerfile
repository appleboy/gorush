FROM alpine:3.17

ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Gorush" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates mailcap && \
  rm -rf /var/cache/apk/*

COPY release/${TARGETOS}/${TARGETARCH}/gorush /bin/

EXPOSE 8088 9000
HEALTHCHECK --start-period=1s --interval=10s --timeout=5s \
  CMD ["/bin/gorush", "--ping"]

ENTRYPOINT ["/bin/gorush"]
