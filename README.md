# gorush

A push notification micro server using [Gin](https://github.com/gin-gonic/gin) framework written in Go (Golang).

Forked from [gorush](https://github.com/appleboy/gorush)

[![GoDoc](https://godoc.org/github.com/appleboy/gorush?status.svg)](https://godoc.org/github.com/appleboy/gorush)
[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/gorush/status.svg)](http://drone.wu-boy.com/appleboy/gorush)
[![codecov](https://codecov.io/gh/appleboy/gorush/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/gorush)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/gorush)](https://goreportcard.com/report/github.com/appleboy/gorush)
[![codebeat badge](https://codebeat.co/badges/0a4eff2d-c9ac-46ed-8fd7-b59942983390)](https://codebeat.co/projects/github-com-appleboy-gorush)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/gorush.svg)](https://hub.docker.com/r/appleboy/gorush/)
[![](https://images.microbadger.com/badges/image/appleboy/gorush.svg)](https://microbadger.com/images/appleboy/gorush "Get your own image badge on microbadger.com")
[![Release](https://github-release-version.herokuapp.com/github/appleboy/gorush/release.svg?style=flat)](https://github.com/appleboy/gorush/releases/latest)

## Contents

- [Support Platform](#support-platform)
- [Features](#features)
- [Memory Usage](#memory-usage)
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
  - [Android Example](#android-example)
  - [Response body](#response-body)
- [Run gRPC service](#run-grpc-service)
- [Run gorush in Docker](#run-gorush-in-docker)
- [Run gorush in Kubernetes](#run-gorush-in-kubernetes)
- [License](#license)

## Support Platform

* [APNS](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/APNSOverview.html)
* [FCM](https://firebase.google.com/)
* [Push API](https://w3c.github.io/push-api/)

## Features

* Support [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging) using [go-fcm](https://github.com/appleboy/go-fcm) library for Android.
* Support [HTTP/2](https://http2.github.io/) Apple Push Notification Service using [apns2](https://github.com/sideshow/apns2) library.
* Support [Push API](https://w3c.github.io/push-api/) using [gowebpush](https://github.com/martijnc/gowebpush) package
* Support [YAML](https://github.com/go-yaml/yaml) configuration.
* Support command line to send single Android or iOS notification.
* Support Web API to send push notification.
* Support graceful restart & zero downtime deploy using [facebook grace](https://github.com/facebookgo/grace).
* Support [HTTP/2](https://http2.github.io/) or HTTP/1.1 protocol.
* Support notification queue and multiple workers.
* Support `/api/stat/app` show notification success and failure counts.
* Support `/api/config` show your [YAML](https://en.wikipedia.org/wiki/YAML) config.
* Support store app stat to memory, [Redis](http://redis.io/), [BoltDB](https://github.com/boltdb/bolt), [BuntDB](https://github.com/tidwall/buntdb) or [LevelDB](https://github.com/syndtr/goleveldb).
* Support `p12` or `pem` formtat of iOS certificate file.
* Support `/sys/stats` show response time, status code count, etc.
* Support for HTTP proxy to Google server (FCM).
* Support retry send notification if server response is fail.
* Support expose [prometheus](https://prometheus.io/) metrics.
* Support install TLS certificates from [Let's Encrypt](https://letsencrypt.org/) automatically.
* Support send notification through [RPC](https://en.wikipedia.org/wiki/Remote_procedure_call) protocol, we use [gRPC](https://grpc.io/) as default framework.

See the [YAML config example](config/config.yml):

[embedmd]:# (config/config.yml yaml)
```yaml
core:
  enabled: true # enabale httpd server
  address: "" # ip address to bind (default: any)
  port: "8088" # ignore this port number if auto_tls is enabled (listen 443).
  worker_num: 0 # default worker number is runtime.NumCPU()
  queue_num: 0 # default queue number is 8192
  max_notification: 100
  sync: false # set true if you need get error message from fail push notification in API response.
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"
  http_proxy: "" # only working for FCM server
  pid:
    enabled: false
    path: "gorush.pid"
    override: true
  auto_tls:
    enabled: false # Automatically install TLS certificates from Let's Encrypt.
    folder: ".cache" # folder for storing TLS certificates
    host: "" # which domains the Let's Encrypt will attempt

grpc:
  enabled: false # enabale gRPC server
  port: 50051

api:
  push_uri: "/api/push"
  stat_go_uri: "/api/stat/go"
  stat_app_uri: "/api/stat/app"
  config_uri: "/api/config"
  sys_stat_uri: "/sys/stats"
  metric_uri: "/metrics"

android:
  enabled: true
  apikey: "YOUR_API_KEY"
  max_retry: 0 # resend fail notification, default value zero is disabled

ios:
  enabled: false
  key_path: "key.pem"
  password: "" # certificate password, default as empty string.
  production: false
  voip_enabled: false
  voip_key_path: "voipkey.pem"
  voip_password: ""
  voip_production: false
  max_retry: 0 # resend fail notification, default value zero is disabled

web:
  enabled: false
  apikey: "YOUR_GCM_API_KEY"
  max_retry: 0 # resend fail notification, default value zero is disabled

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console, or define log path like "log/error_log"
  error_level: "error"
  hide_token: true

stat:
  engine: "memory" # support memory, redis, boltdb, buntdb or leveldb
  redis:
    addr: "localhost:6379"
    password: ""
    db: 0
  boltdb:
    path: "bolt.db"
    bucket: "gorush"
  buntdb:
    path: "bunt.db"
  leveldb:
    path: "level.db"
```

## Memory Usage

Memory average usage: **28Mb** (the total bytes of memory obtained from the OS.)

![memory usage](screenshot/memory.png)

Test Command:

```sh
$ for i in {1..9999999}; do bat -b.N=1000 -b.C=100 POST localhost:8088/api/push notifications:=@notification.json; sleep 1;  done
```

## Basic Usage

How to send push notification using `gorush` command? (Android or iOS)

### Download a binary

The pre-compiled binaries can be downloaded from [release page](https://github.com/appleboy/gorush/releases).

With `Go` installed

```
$ go get -u -v github.com/appleboy/gorush
```

On linux

```sh
$ wget https://github.com/appleboy/gorush/releases/download/v1.9.0/gorush-v1.9.0-linux-amd64 -O gorush
```

On OS X

```sh
$ wget https://github.com/appleboy/gorush/releases/download/v1.9.0/gorush-v1.9.0-darwin-amd64 -O gorush
```

On Windows

```sh
$ wget https://github.com/appleboy/gorush/releases/download/v1.9.0/gorush-v1.9.0-windows-amd64.exe -O gorush.exe
```

### Command Usage

```
  ________                              .__
 /  _____/   ____ _______  __ __  ______|  |__
/   \  ___  /  _ \\_  __ \|  |  \/  ___/|  |  \
\    \_\  \(  <_> )|  | \/|  |  /\___ \ |   Y  \
 \______  / \____/ |__|   |____//____  >|___|  /
        \/                           \/      \/

Usage: gorush [options]

Server Options:
    -p, --port <port>                Use port for clients (default: 8088)
    -c, --config <file>              Configuration file
    -m, --message <message>          Notification message
    -t, --token <token>              Notification token
    --title <title>                  Notification title
    --proxy <proxy>                  Proxy URL (only for FCM)
    --pid <pid path>                 Process identifier path
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
* `-k`: [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging) api key
* `-t`: Device token.
* `--title`: Notification title.
* `--proxy`: Set http proxy url. (only working for FCM)

### Send iOS notification

Send single notification with the following command.

```bash
$ gorush -ios -m="your message" -i="your certificate path" -t="device token" -topic="apns topic"
```

* `-m`: Notification message.
* `-i`: Apple Push Notification Certificate path (`pem` or `p12` file).
* `-t`: Device token.
* `--title`: Notification title.
* `--topic`: The topic of the remote notification.
* `--password`: The certificate password.

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
$ http -v --verify=no --json GET http://localhost:8088/api/stat/go
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
  "version": "v1.6.2",
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

### GET /metrics

Support expose [prometheus](https://prometheus.io/) metrics.

![metrics screenshot](screenshot/metrics.png)

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

Request body is a JSON containing the following fields.

| name          | type         | description                                                     | required | note                                                                              |
|---------------|--------------|-----------------------------------------------------------------|----------|-----------------------------------------------------------------------------------|
| notifications | object array | Notifications that are being sent                               | o        |                                                                                   |
| sync          | bool         | Wait for push services responses before replying to the request | -        | Error messages are only returned when this is true. Default value is config based |

The following is a parameter table for each notification in the notifications array.

Request body must has a notifications array. The following is a parameter table for each notification.

| name                    | type         | description                                                                                       | required | note                                                          |
|-------------------------|--------------|---------------------------------------------------------------------------------------------------|----------|---------------------------------------------------------------|
| tokens                  | string array | device tokens                                                                                     | o        |                                                               |
| subscriptions           | object array | Web Push subscription objects                                                                     | o        | Required if platform is 3                                     |
| platform                | int          | platform(iOS,Android,Web)                                                                         | o        | 1=iOS, 2=Android, 3=Web                                       |
| message                 | string       | message for notification                                                                          | -        |                                                               |
| title                   | string       | notification title                                                                                | -        |                                                               |
| priority                | string       | Sets the priority of the message.                                                                 | -        | `normal` or `high`                                            |
| content_available       | bool         | data messages wake the app by default.                                                            | -        |                                                               |
| sound                   | string       | sound type                                                                                        | -        |                                                               |
| data                    | string array | extensible partition                                                                              | -        | payload for Web is taken from this field                      |
| retry                   | int          | retry send notification if fail response from server. Value must be small than `max_retry` field. | -        |                                                               |
| api_key                 | string       | Android api key                                                                                   | -        | only Android                                                  |
| to                      | string       | The value must be a registration token, notification key, or topic.                               | -        | only Android                                                  |
| collapse_key            | string       | a key for collapsing notifications                                                                | -        | only Android                                                  |
| delay_while_idle        | bool         | a flag for device idling                                                                          | -        | only Android                                                  |
| time_to_live            | uint         | expiration of message kept on FCM storage                                                         | -        | only Android                                                  |
| restricted_package_name | string       | the package name of the application                                                               | -        | only Android                                                  |
| dry_run                 | bool         | allows developers to test a request without actually sending a message                            | -        | only Android                                                  |
| notification            | string array | payload of a FCM message                                                                          | -        | only Android. See the [detail](#android-notification-payload) |
| expiration              | int          | expiration for notification                                                                       | -        | only iOS                                                      |
| apns_id                 | string       | A canonical UUID that identifies the notification                                                 | -        | only iOS                                                      |
| topic                   | string       | topic of the remote notification                                                                  | -        | only iOS                                                      |
| badge                   | int          | badge count                                                                                       | -        | only iOS                                                      |
| category                | string       | the UIMutableUserNotificationCategory object                                                      | -        | only iOS                                                      |
| alert                   | string array | payload of a iOS message                                                                          | -        | only iOS. See the [detail](#ios-alert-payload)                |
| mutable_content         | bool         | enable Notification Service app extension.                                                        | -        | only iOS(10.0+).                                              |
| voip                    | bool         | send VoIP push instead of normal push.                                                            | -        | only iOS.                                                     |

### iOS alert payload

| name           | type             | description                                                                                      | required | note |
|----------------|------------------|--------------------------------------------------------------------------------------------------|----------|------|
| title          | string           | Apple Watch & Safari display this string as part of the notification interface.                  | -        |      |
| body           | string           | The text of the alert message.                  | -        |      |
| subtitle       | string           | Apple Watch & Safari display this string as part of the notification interface.                  | -        |      |
| action         | string           | The label of the action button. This one is required for Safari Push Notifications.              | -        |      |
| action-loc-key | string           | If a string is specified, the system displays an alert that includes the Close and View buttons. | -        |      |
| launch-image   | string           | The filename of an image file in the app bundle, with or without the filename extension.         | -        |      |
| loc-args       | array of strings | Variable string values to appear in place of the format specifiers in loc-key.                   | -        |      |
| loc-key        | string           | A key to an alert-message string in a Localizable.strings file for the current localization.     | -        |      |
| title-loc-args | array of strings | Variable string values to appear in place of the format specifiers in title-loc-key.             | -        |      |
| title-loc-key  | string           | The key to a title string in the Localizable.strings file for the current localization.          | -        |      |

See more detail about [APNs Remote Notification Payload](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html).

### Android notification payload

| name           | type   | description                                                                                               | required | note |
|----------------|--------|-----------------------------------------------------------------------------------------------------------|----------|------|
| icon           | string | Indicates notification icon.                                                                              | -        |      |
| tag            | string | Indicates whether each notification message results in a new entry on the notification center on Android. | -        |      |
| color          | string | Indicates color of the icon, expressed in #rrggbb format                                                  | -        |      |
| click_action   | string | The action associated with a user click on the notification.                                              | -        |      |
| body_loc_key   | string | Indicates the key to the body string for localization.                                                    | -        |      |
| body_loc_args  | string | Indicates the string value to replace format specifiers in body string for localization.                  | -        |      |
| title_loc_key  | string | Indicates the key to the title string for localization.                                                   | -        |      |
| title_loc_args | string | Indicates the string value to replace format specifiers in title string for localization.                 | -        |      |

See more detail about [Firebase Cloud Messaging HTTP Protocol reference](https://firebase.google.com/docs/cloud-messaging/http-server-ref#send-downstream).

### Web Push subscription

| name      | type   | description                                                                | required | note |
|-----------|--------|----------------------------------------------------------------------------|----------|------|
| endpoint  | string | Endpoint URL from PushSubscription browser object.                         | o        |      |
| key       | string | P-256 ECDH Diffie-Hellman public key from PushSubscription browser object. | o        |      |
| auth      | string | Authentication secret from PushSubscription browser object.                | o        |      |

See more detail about [PushSubscription interface](https://w3c.github.io/push-api/#pushsubscription-interface)

### iOS Example

Send normal notification.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello World iOS!"
    }
  ]
```

The following payload asks the system to display an alert with a Close button and a single action button.The title and body keys provide the contents of the alert. The “PLAY” string is used to retrieve a localized string from the appropriate Localizable.strings file of the app. The resulting string is used by the alert as the title of an action button. This payload also asks the system to badge the app’s icon with the number 5.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "badge": 5,
      "alert": {
        "title" : "Game Request",
        "body" : "Bob wants to play poker",
        "action-loc-key" : "PLAY"
      }
    }
  ]
```

The following payload specifies that the device should display an alert message, plays a sound, and badges the app’s icon.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "You got your emails.",
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
      "data": {
        "key1": "welcome",
        "key2": 2
      }
    }
  ]
```

Send VoIP push.

```json
  "notifications": [
    {
      "tokens": ["token_a", "token_b"],
      "platform": 1,
      "message": "Hello VoIP iOS!",
      "voip": true,
      "data": {
        "key1": "Incoming call",
        "key2": "Peter"
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

### Web Example

Send normal notification.

```json
  "notifications": [
    {
      "subscriptions": [{
        "endpoint": "endpoint_a",
        "key": "key_a",
        "auth": "auth_a"
      },{
        "endpoint": "endpoint_b",
        "key": "key_b",
        "auth": "auth_b"
      }],
      "platform": 3,
      "data": {
        "message": "Hello World Web!",
        "title": "You got message"
      }
    }
  ]
```

### Response body

Error response message table:

| status code | message                                    |
|-------------|--------------------------------------------|
| 400         | Missing `notifications` field.             |
| 400         | Notifications field is empty.              |
| 400         | Number of notifications(50) over limit(10) |

Success response on async request:

```json
{
  "counts": 60,
  "logs": [],
  "success": "ok"
}
```

Success response on sync request:

```json
{
  "apnsFailedResults": {},
  "gcmFailedResults": {},
  "success": "ok",
  "webFailedResults": {}
}

If you need error logs from sending fail notifications, please set `sync` as `true` on yaml config.

```diff
core:
  port: "8088" # ignore this port number if auto_tls is enabled (listen 443).
  worker_num: 0 # default worker number is runtime.NumCPU()
  queue_num: 0 # default queue number is 8192
  max_notification: 100
- sync: false 
+ sync: true
```

See the following error format.

```json
{
  "counts": 60,
  "logs": [
    {
      "type": "failed-push",
      "platform": "android",
      "token": "*******",
      "message": "Hello World Android!",
      "error": "InvalidRegistration"
    },
    {
      "type": "failed-push",
      "platform": "ios",
      "token": "*****",
      "message": "Hello World iOS1111!",
      "error": "Post https://api.push.apple.com/3/device/bbbbb: remote error: tls: revoked certificate"
    },
    {
      "type": "failed-push",
      "platform": "ios",
      "token": "*******",
      "message": "Hello World iOS222!",
      "error": "Post https://api.push.apple.com/3/device/token_b: remote error: tls: revoked certificate"
    }
  ],
  "success": "ok"
}
```

Response with some fails for each platform on sync request:

```json
{
  "apnsFailedResults": {
    "apns_token_a": {
      "StatusCode": 400,
      "Reason": "BadDeviceToken",
      "ApnsID": "apns_id_a",
      "Timestamp": "0001-01-01T00:00:00Z"
    }
  },
  "gcmFailedResults": {
    "gcm_token_a": "InvalidRegistration"
  },
  "success": "ok",
  "webFailedResults": {
    "chrome_endpoint_a": {
      "StatusCode": 400,
      "Body": "<HTML>\n<HEAD>\n<TITLE>UnauthorizedRegistration</TITLE>\n</HEAD>\n<BODY BGCOLOR=\"#FFFFFF\" TEXT=\"#000000\">\n<H1>UnauthorizedRegistration</H1>\n<H2>Error 400</H2>\n</BODY>\n</HTML>\n"
    },
    "firefox_endpoint_a": {
      "StatusCode": 410,
      "Body": "{\"code\": 410, \"errno\": 106, \"error\": \"\", \"more_info\": \"http://autopush.readthedocs.io/en/latest/http.html#error-codes\", \"message\": \"Request did not validate No such subscription\"}"
    },
    "firefox_endpoint_b": {
      "StatusCode": 404,
      "Body": "{\"code\": 404, \"errno\": 102, \"error\": \"Not Found\", \"more_info\": \"http://autopush.readthedocs.io/en/latest/http.html#error-codes\", \"message\": \"Request did not validate invalid token\"}"
    }
  }
}
```


## Run gRPC service

Gorush support [gRPC](https://grpc.io/) service. You can enable the gRPC in `config.yml`, default as disabled. The following example code for Golang to send single notification.

[embedmd]:# (rpc/example/go/client.go go)
```go
package main

import (
	"log"

	pb "github.com/appleboy/gorush/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGorushClient(conn)

	r, err := c.Send(context.Background(), &pb.NotificationRequest{
		Platform: 2,
		Tokens:   []string{"1234567890"},
		Message:  "test message",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Success: %t\n", r.Success)
	log.Printf("Count: %d\n", r.Counts)
}
```

See the Node.js example and see more detail frome [README](rpc/example/node/README.md): 

[embedmd]:# (rpc/example/node/client.js js)
```js
var messages = require('./gorush_pb');
var services = require('./gorush_grpc_pb');

var grpc = require('grpc');

function main() {
  var client = new services.GorushClient('localhost:50051',
    grpc.credentials.createInsecure());
  var request = new messages.NotificationRequest();
  request.setPlatform(2);
  request.setTokensList(["1234567890"]);
  request.setMessage("Hello!!");
  client.send(request, function (err, response) {
    console.log('Success:', response.getSuccess());
    console.log('Counts:', response.getCounts());
  });
}

main();
```

## Run gorush in Docker

Set up `gorush` in the cloud in under 5 minutes with zero knowledge of Golang or Linux shell using our [gorush Docker image](https://hub.docker.com/r/appleboy/gorush/).

```bash
$ docker pull appleboy/gorush
$ docker run --name gorush -p 80:8088 appleboy/gorush
```

Run `gorush` with your own config file.

```bash
$ docker pull appleboy/gorush
$ docker run --name gorush -v ${PWD}/config.yml:/config.yml -p 80:8088 appleboy/gorush
```

Testing your gorush server using [httpie](https://github.com/jkbrzt/httpie) command.

```bash
$ http -v --verify=no --json GET http://your.docker.host/api/stat/go
```

![statue screenshot](screenshot/status.png)

## Run gorush in Kubernetes

Please make sure you are install [Minikube](https://kubernetes.io/docs/tutorials/stateless-application/hello-minikube/) first.

### Create a Minikube cluster

```sh
$ minikube start --vm-driver=xhyve
```

### Quick Start

Start the gorush with one command:

```sh
$ kubectl create -f k8s
deployment "frontend" created
service "frontend" created
deployment "redis-master" created
service "redis-master" created
```

Then, list all your Services:

```sh
$ kubectl get services
NAME           CLUSTER-IP   EXTERNAL-IP   PORT(S)          AGE
frontend       10.0.0.156   <pending>     8088:32517/TCP   30s
kubernetes     10.0.0.1     <none>        443/TCP          14d
redis-master   10.0.0.67    <none>        6379/TCP         30s
```

### view the gorush home page

```sh
$ minikube service frontend
```

### Clean up the gorush:

```sh
$ kubectl delete -f k8s
```

## License

Copyright 2016 Bo-Yi Wu [@appleboy](https://twitter.com/appleboy).

Licensed under the MIT License.
