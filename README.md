# Gopush

A push notification server using [Gin](https://github.com/gin-gonic/gin) framework written in Go (Golang).

[![Build Status](https://travis-ci.org/appleboy/gofight.svg?branch=master)](https://travis-ci.org/appleboy/gofight) [![Coverage Status](https://coveralls.io/repos/github/appleboy/gopush/badge.svg?branch=master)](https://coveralls.io/github/appleboy/gopush?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/gopush)](https://goreportcard.com/report/github.com/appleboy/gopush) [![codebeat badge](https://codebeat.co/badges/ee01d852-b5e8-465a-ad93-631d738818ff)](https://codebeat.co/projects/github-com-appleboy-gopush)

## Support Platform

* [APNS](https://developer.apple.com/library/ios/documentation/networkinginternet/conceptual/remotenotificationspg/Chapters/ApplePushService.html)
* [GCM](https://developer.android.com/google/gcm/index.html)

## Feature

* Support [Google Cloud Message](https://developers.google.com/cloud-messaging/) using [go-gcm](https://github.com/google/go-gcm) library for Android.
* Support [HTTP/2](https://http2.github.io/) Apple Push Notification Service using [apns2](https://github.com/sideshow/apns2) library.
* Support [YAML](https://github.com/go-yaml/yaml) configuration.
* Support command line to send single Android or iOS notification.
* Support Web API to send push notification.
* Support zero downtime restarts for go servers using [endless](https://github.com/fvbock/endless).
* Support [HTTP/2](https://http2.github.io/) or HTTP/1.1 protocol.

See the [YAML config example](config/config.yml):

```yaml
core:
  port: "8088"
  max_notification: 100
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"

api:
  push_uri: "/api/push"
  stat_go_uri: "/api/status"

android:
  enabled: true
  apikey: "YOUR_API_KEY"

ios:
  enabled: false
  pem_cert_path: "cert.pem"
  pem_key_path: "key.pem"
  production: false

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console, or define log path like "log/error_log"
  error_level: "error"
```

## Basic Usage

How to send push notification using `gopush` command? (Android or iOS)

### Download a binary

The pre-compiled binaries can be downloaded from [release page](https://github.com/appleboy/gopush/releases).

### Send Android notification

Send single notification with the following command.

```bash
$ gopush -android -m="your message" -k="API Key" -t="Device token"
```

* `-m`: Notification message.
* `-k`: [Google cloud message](https://developers.google.com/cloud-messaging/) api key
* `-t`: Device token.

### Send iOS notification

Send single notification with the following command.

```bash
$ gopush -ios -m="your message" -i="API Key" -t="Device token"
```

* `-m`: Notification message.
* `-i`: Apple Push Notification Certificate path (`pem` file).
* `-t`: Device token.

The default endpoint is APNs development. Please add `-production` flag for APNs production push endpoint.

```bash
$ gopush -ios -m="your message" -i="API Key" -t="Device token" -production
```

## Run gopush web server

Please make sure your [config.yml](config/config.yml) exist. Default port is `8088`.

```bash
$ gopush -c config.yml
```

Test status of api server using [httpie](https://github.com/jkbrzt/httpie) tool:

```bash
$ http -v --verify=no --json GET https://localhost:8088/api/status
```

![statue screenshot](screenshot/status.png)

## Run gopush in Docker

Set up `gopush` in the cloud in under 5 minutes with zero knowledge of Golang or Linux shell using our [gopush Docker image](https://hub.docker.com/r/appleboy/gopush/).

```bash
$ docker pull appleboy/gopush
$ docker run -name gopush -p 80:8088 appleboy/gopush
```

Testing your gopush server.

```bash
$ http -v --verify=no --json GET http://your.docker.host/api/status
```

## Web API

Gopush support the following API.

* **GET**  `/api/status` Golang cpu, memory, gc, etc information. Thanks for [golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler).
* **POST** `/api/push` push ios and android notifications.

### POST /api/push

Simple send iOS notification example, the `platform` value is `1`:

```json
{
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello World iOS!"
    }
  ]
}
```

Simple send Android notification example, the `platform` value is `2`:

```json
{
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 2,
      "message": "Hello World Android!"
    }
  ]
}
```

Send multiple notifications as below:

```json
{
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello World iOS!"
    },
    {
      "tokens": ["token_a", "token_b"],
      "platform": 2,
      "message": "Hello World Android!"
    },
    {
      "tokens": ["token_a", "token_b"],
      "platform": 2,
      "message": "Hello World!"
    },
    .....
  ]
}
```

Request body must has a notifications array. The following is a parameter table for each notification.

|name|type|description|required|note|
|-------|-------|--------|--------|---------|
|tokens|string array|device tokens|o||
|platform|int|platform(iOS,Android)|o|1=iOS, 2=Android|
|message|string|message for notification|o||
|priority|string|Sets the priority of the message.|-||
|content_available|bool|data messages wake the app by default.|-||
|api_key|string|Android api key|-|only Android|
|to|string|The value must be a registration token, notification key, or topic.|-|only Android|
|collapse_key|string|a key for collapsing notifications|-|only Android|
|delay_while_idle|bool|a flag for device idling|-|only Android|
|time_to_live|int|expiration of message kept on GCM storage|-|only Android|
|restricted_package_name|string|the package name of the application|-|only Android|
|dry_run|bool|allows developers to test a request without actually sending a message|-|only Android|
|data|string array|data payload of a GCM message|-|only Android|
|notification|string array|payload of a GCM message|-|only Android|
|expiration|int|expiration for notification|-|only iOS|
|apns_id|string|A canonical UUID that identifies the notification|-|only iOS|
|topic|string|topic of the remote notification|-|only iOS|
|badge|int|badge count|-|only iOS|
|sound|string|sound type|-|only iOS|
|category|string|the UIMutableUserNotificationCategory object|-|only iOS|
|extend|string array|extensible partition|-|only iOS|
|alert|string array|payload of a iOS message|-|only iOS|

## License

Copyright 2016 Bo-Yi Wu [@appleboy](https://twitter.com/appleboy).

Licensed under the MIT License.
