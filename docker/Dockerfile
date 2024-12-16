FROM alpine:3.21

ARG TARGETOS
ARG TARGETARCH
ARG USER=gorush
ENV HOME /home/$USER

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Gorush" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

# add new user
RUN adduser -D $USER
RUN apk add --no-cache ca-certificates mailcap && \
  rm -rf /var/cache/apk/*

COPY release/${TARGETOS}/${TARGETARCH}/gorush /bin/

USER $USER
WORKDIR $HOME

EXPOSE 8088 9000
HEALTHCHECK --start-period=1s --interval=10s --timeout=5s \
  CMD ["/bin/gorush", "--ping"]

ENTRYPOINT ["/bin/gorush"]
