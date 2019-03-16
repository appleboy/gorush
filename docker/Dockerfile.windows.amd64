FROM microsoft/nanoserver:10.0.14393.1884

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Gorush" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

COPY release/gorush.exe C:/bin/gorush.exe

EXPOSE 8088 9000
HEALTHCHECK --start-period=2s --interval=10s --timeout=5s \
  CMD ["\\gorush.exe", "--ping"]

ENTRYPOINT [ "C:\\bin\\gorush.exe" ]
