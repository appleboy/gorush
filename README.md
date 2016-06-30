# gorush

A push notification micro server using [Gin](https://github.com/gin-gonic/gin) framework written in Go (Golang).

[![GoDoc](https://godoc.org/github.com/appleboy/gorush?status.svg)](https://godoc.org/github.com/appleboy/gorush) [![Build Status](https://travis-ci.org/appleboy/gorush.svg?branch=master)](https://travis-ci.org/appleboy/gorush) [![Coverage Status](https://coveralls.io/repos/github/appleboy/gorush/badge.svg?branch=master)](https://coveralls.io/github/appleboy/gorush?branch=master) [![codecov](https://codecov.io/gh/appleboy/gorush/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/gorush) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/gorush)](https://goreportcard.com/report/github.com/appleboy/gorush) [![codebeat badge](https://codebeat.co/badges/0a4eff2d-c9ac-46ed-8fd7-b59942983390)](https://codebeat.co/projects/github-com-appleboy-gorush)

## Contents

- [Support Platform](#support-platform)
- [Features](#features)
- [Basic Usage](#basic-usage)
  - [Download a binary](#download-a-binary)
  - [Command Usage](#command-usage)
  - [Send Android notification](#send-android-notification)
  - [Send iOS notification](#send-ios-notification)
- [Run gorush web server](#run-gorush-web-server)
- [Web API](#web-api)
  - [GET /api/stat/go](#get-apistatgo)
  - [GET /api/stat/app](#get-apistatapp)
  - [GET /sys/stats](#get-sysstats)
  - [POST /api/push](#post-apipush)
  - [Request body](#request-body)
  - [iOS alert payload](#ios-alert-payload)
  - [Android notification payload](#android-notification-payload)
  - [iOS Example](#ios-example)
  - [Android Example](#Android-Example)
  - [Response body](#Response-body)
- [Run gorush in Docker](#run-gorush-in-docker)
- [License](#license)

## Support Platform

* [APNS](https://developer.apple.com/library/ios/documentation/networkinginternet/conceptual/remotenotificationspg/Chapters/ApplePushService.html)
* [GCM](https://developer.android.com/google/gcm/index.html)

## Features

* Support [Google Cloud Message](https://developers.google.com/cloud-messaging/) using [go-gcm](https://github.com/google/go-gcm) library for Android.
* Support [HTTP/2](https://http2.github.io/) Apple Push Notification Service using [apns2](https://github.com/sideshow/apns2) library.
* Support [YAML](https://github.com/go-yaml/yaml) configuration.
* Support command line to send single Android or iOS notification.
* Support Web API to send push notification.
* Support zero downtime restarts for go servers using [endless](https://github.com/fvbock/endless).
* Support [HTTP/2](https://http2.github.io/) or HTTP/1.1 protocol.
* Support notification queue and multiple workers.
* Support `/api/stat/app` show notification success and failure counts.
* Support `/api/config` show your [YAML](https://en.wikipedia.org/wiki/YAML) config.
* Support store app stat to memory, [Redis](http://redis.io/) or [BoltDB](https://github.com/boltdb/bolt).
* Support `p12` or `pem` formtat of iOS certificate file.
* Support `/sys/stats` show response time, status code count, etc.

See the [YAML config example](config/config.yml):

```yaml
core:
  port: "8088"
  worker_num: 8
  queue_num: 8192
  max_notification: 100
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"

api:
  push_uri: "/api/push"
  stat_go_uri: "/api/stat/go"
  stat_app_uri: "/api/stat/app"
  config_uri: "/api/config"
  sys_stat_uri: "/sys/stats"

android:
  enabled: true
  apikey: "YOUR_API_KEY"

ios:
  enabled: false
  key_path: "key.pem"
  password: "" # certificate password, default as empty string.
  production: false

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console, or define log path like "log/error_log"
  error_level: "error"
  hide_token: true

stat:
  engine: "memory" # support memory, redis or boltdb
  redis:
    addr: "localhost:6379"
    password: ""
    db: 0
  boltdb:
    path: "gorush.db"
    bucket: "gorush"
```

## Basic Usage

How to send push notification using `gorush` command? (Android or iOS)

### Download a binary

The pre-compiled binaries can be downloaded from [release page](https://github.com/appleboy/gorush/releases).

### Command Usage

```
Usage: gorush [options]

Server Options:
    -p, --port <port>                Use port for clients (default: 8088)
    -c, --config <file>              Configuration file
    -m, --message <message>          Notification message
    -t, --token <token>              Notification token
iOS Options:
    -i, --key <file>                 certificate key file path
    -P, --password <password>        certificate key password
    --topic <topic>                  iOS topic
    --ios                            enabled iOS (default: false)
    --production                     iOS production mode (default: false)
Android Options:
    -k, --apikey <api_key>           Android API Key
    --android                        enabled android (default: false)
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
```

### Send Android notification

Send single notification with the following command.

```bash
$ gorush -android -m="your message" -k="API Key" -t="Device token"
```

* `-m`: Notification message.
* `-k`: [Google cloud message](https://developers.google.com/cloud-messaging/) api key
* `-t`: Device token.

### Send iOS notification

Send single notification with the following command.

```bash
$ gorush -ios -m="your message" -i="your certificate path" -t="device token" -topic="apns topic"
```

* `-m`: Notification message.
* `-i`: Apple Push Notification Certificate path (`pem` or `p12` file).
* `-t`: Device token.
* `-topic`: The topic of the remote notification.
* `-password`: The certificate password.

The default endpoint is APNs development. Please add `-production` flag for APNs production push endpoint.

```bash
$ gorush -ios -m="your message" -i="your certificate path" -t="device token" -production
```

## Run gorush web server

Please make sure your [config.yml](config/config.yml) exist. Default port is `8088`.

```bash
$ gorush -c config.yml
```

Get go status of api server using [httpie](https://github.com/jkbrzt/httpie) tool:

```bash
$ http -v --verify=no --json GET https://localhost:8088/api/stat/go
```

## Web API

Gorush support the following API.

* **GET**  `/api/stat/go` Golang cpu, memory, gc, etc information. Thanks for [golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler).
* **GET**  `/api/stat/app` show notification success and failure counts.
* **GET**  `/api/config` show server yml config file.
* **POST** `/api/push` push ios and android notifications.

### GET /api/stat/go

Golang cpu, memory, gc, etc information. Response with `200` http status code.

```json
{
  "time": 1460686815848046600,
  "go_version": "go1.6.1",
  "go_os": "darwin",
  "go_arch": "amd64",
  "cpu_num": 4,
  "goroutine_num": 15,
  "gomaxprocs": 4,
  "cgo_call_num": 1,
  "memory_alloc": 7455192,
  "memory_total_alloc": 8935464,
  "memory_sys": 12560632,
  "memory_lookups": 17,
  "memory_mallocs": 31426,
  "memory_frees": 11772,
  "memory_stack": 524288,
  "heap_alloc": 7455192,
  "heap_sys": 8912896,
  "heap_idle": 909312,
  "heap_inuse": 8003584,
  "heap_released": 0,
  "heap_objects": 19654,
  "gc_next": 9754725,
  "gc_last": 1460686815762559700,
  "gc_num": 2,
  "gc_per_second": 0,
  "gc_pause_per_second": 0,
  "gc_pause": [
    0.326576,
    0.227096
  ]
}
```

### GET /api/stat/app

Show success or failure counts information of notification.

```json
{
  "queue_max": 8192,
  "queue_usage": 0,
  "total_count": 77,
  "ios": {
    "push_success": 19,
    "push_error": 38
  },
  "android": {
    "push_success": 10,
    "push_error": 10
  }
}
```

### GET /sys/stats

Show response time, status code count, etc.

```json
{
  "pid": 80332,
  "uptime": "1m42.428010614s",
  "uptime_sec": 102.428010614,
  "time": "2016-06-26 12:27:11.675973571 +0800 CST",
  "unixtime": 1466915231,
  "status_code_count": { },
  "total_status_code_count": {
    "200": 5
  },
  "count": 0,
  "total_count": 5,
  "total_response_time": "10.422441ms",
  "total_response_time_sec": 0.010422441000000001,
  "average_response_time": "2.084488ms",
  "average_response_time_sec": 0.0020844880000000002
}
```

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

See more example about [iOS](#ios-example) or [Android](#android-example).

### Request body

Request body must has a notifications array. The following is a parameter table for each notification.

|name|type|description|required|note|
|-------|-------|--------|--------|---------|
|tokens|string array|device tokens|o||
|platform|int|platform(iOS,Android)|o|1=iOS, 2=Android|
|message|string|message for notification|o||
|title|string|notification title|-||
|priority|string|Sets the priority of the message.|-|`normal` or `high`|
|content_available|bool|data messages wake the app by default.|-||
|sound|string|sound type|-||
|data|string array|extensible partition|-||
|api_key|string|Android api key|-|only Android|
|to|string|The value must be a registration token, notification key, or topic.|-|only Android|
|collapse_key|string|a key for collapsing notifications|-|only Android|
|delay_while_idle|bool|a flag for device idling|-|only Android|
|time_to_live|uint|expiration of message kept on GCM storage|-|only Android|
|restricted_package_name|string|the package name of the application|-|only Android|
|dry_run|bool|allows developers to test a request without actually sending a message|-|only Android|
|notification|string array|payload of a GCM message|-|only Android. See the [detail](#android-notification-payload)|
|expiration|int|expiration for notification|-|only iOS|
|apns_id|string|A canonical UUID that identifies the notification|-|only iOS|
|topic|string|topic of the remote notification|-|only iOS|
|badge|int|badge count|-|only iOS|
|category|string|the UIMutableUserNotificationCategory object|-|only iOS|
|alert|string array|payload of a iOS message|-|only iOS. See the [detail](#ios-alert-payload)|

### iOS alert payload

|name|type|description|required|note|
|-------|-------|--------|--------|---------|
|action|string|The label of the action button. This one is required for Safari Push Notifications.|-||
|action-loc-key|string|If a string is specified, the system displays an alert that includes the Close and View buttons.|-||
|launch-image|string|The filename of an image file in the app bundle, with or without the filename extension.|-||
|loc-args|array of strings|Variable string values to appear in place of the format specifiers in loc-key.|-||
|loc-key|string|A key to an alert-message string in a Localizable.strings file for the current localization.|-||
|title-loc-args|array of strings|Variable string values to appear in place of the format specifiers in title-loc-key.|-||
|title-loc-key|string|The key to a title string in the Localizable.strings file for the current localization.|-||

See more detail about [APNs Remote Notification Payload](https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/TheNotificationPayload.html#//apple_ref/doc/uid/TP40008194-CH107-SW1).

### Android notification payload

|name|type|description|required|note|
|-------|-------|--------|--------|---------|
|icon|string|Indicates notification icon.|-||
|tag|string|Indicates whether each notification message results in a new entry on the notification center on Android.|-||
|color|string|Indicates color of the icon, expressed in #rrggbb format|-||
|click_action|string|The action associated with a user click on the notification.|-||
|body_loc_key|string|Indicates the key to the body string for localization.|-||
|body_loc_args|string|Indicates the string value to replace format specifiers in body string for localization.|-||
|title_loc_key|string|Indicates the key to the title string for localization.|-||
|title_loc_args|string|Indicates the string value to replace format specifiers in title string for localization.|-||

See more detail about [GCM server reference](https://developers.google.com/cloud-messaging/http-server-ref#send-downstream).

### iOS Example

Send normal notification.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello World iOS!",
      "title": "You got message"
    }
  ]
```

The app icon be badged with the number `9` and that a bundled alert sound be played when the notification is delivered.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello World iOS!",
      "title": "You got message",
      "badge": 9,
      "sound": "bingbong.aiff"
    }
  ]
```

Add other fields which user defined via `data` field.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello World iOS!",
      "title": "You got message",
      "data": {
        "key1": "welcome",
        "key2": 2
      }
    }
  ]
```

### Android Example

Send normal notification.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 2,
      "message": "Hello World Android!",
      "title": "You got message"
    }
  ]
```

Add `notification` payload.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 2,
      "message": "Hello World Android!",
      "title": "You got message",
      "notification" : {
        "icon": "myicon",
        "color": "#112244"
      }
    }
  ]
```

Add other fields which user defined via `data` field.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 2,
      "message": "Hello World Android!",
      "title": "You got message",
      "data": {
       "Nick" : "Mario",
       "body" : "great match!",
       "Room" : "PortugalVSDenmark"
      }
    }
  ]
```

### Response body

Error response message table:

|status code|message|
|-------|-------|
|400|Missing `notifications` field.|
|400|Notifications field is empty.|
|400|Number of notifications(50) over limit(10)|

Success response:

```json
{
  "success": "ok"
}
```

## Run gorush in Docker

Set up `gorush` in the cloud in under 5 minutes with zero knowledge of Golang or Linux shell using our [gorush Docker image](https://hub.docker.com/r/appleboy/gorush/).

```bash
$ docker pull appleboy/gorush
$ docker run --name gorush -p 80:8088 appleboy/gorush
```

Testing your gorush server using [httpie](https://github.com/jkbrzt/httpie) command.

```bash
$ http -v --verify=no --json GET http://your.docker.host/api/stat/go
```

![statue screenshot](screenshot/status.png)

## License

Copyright 2016 Bo-Yi Wu [@appleboy](https://twitter.com/appleboy).

Licensed under the MIT License.
