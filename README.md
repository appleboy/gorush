# Gopush

A push notification server using [Gin](https://github.com/gin-gonic/gin) framework written in Go (Golang).

[![Build Status](https://travis-ci.org/appleboy/gofight.svg?branch=master)](https://travis-ci.org/appleboy/gofight) [![Coverage Status](https://coveralls.io/repos/github/appleboy/gopush/badge.svg?branch=master)](https://coveralls.io/github/appleboy/gopush?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/gopush)](https://goreportcard.com/report/github.com/appleboy/gopush) [![codebeat badge](https://codebeat.co/badges/ee01d852-b5e8-465a-ad93-631d738818ff)](https://codebeat.co/projects/github-com-appleboy-gopush)

## Feature

* Support [Google Cloud Message](https://developers.google.com/cloud-messaging/) using [go-gcm](https://github.com/google/go-gcm) library for Android.
* Support [HTTP/2](https://http2.github.io/) Apple Push Notification Service using [apns2](https://github.com/sideshow/apns2) library.
* Support [YAML](https://github.com/go-yaml/yaml) configuration.

See the [YAML config eample](config/config.yaml):

```yaml
core:
  port: "8088"
  notification_max: 100
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"

api:
  push_uri: "/api/push"
  stat_go_uri: "/api/status"

android:
  enabled: false
  apikey: ""

ios:
  enabled: false
  pem_cert_path: "cert.pem"
  pem_key_path: "key.pem"
  production: false

log:
  format: "string" # string or json
  access_log: "stdout"
  access_level: "debug"
  error_log: "stderr"
  error_level: "error"
```

## License

Copyright 2016 Bo-Yi Wu [@appleboy](https://twitter.com/appleboy).

Licensed under the MIT License.
