# APNS/2

APNS/2 is a go package designed for simple, flexible and fast Apple Push Notifications on iOS, OSX and Safari using the new HTTP/2 Push provider API.

[![Build Status](https://travis-ci.org/sideshow/apns2.svg?branch=master)](https://travis-ci.org/sideshow/apns2)  [![Coverage Status](https://coveralls.io/repos/sideshow/apns2/badge.svg?branch=master&service=github)](https://coveralls.io/github/sideshow/apns2?branch=master)  [![GoDoc](https://godoc.org/github.com/sideshow/apns2?status.svg)](https://godoc.org/github.com/sideshow/apns2)

## Features

- Uses new Apple APNs HTTP/2 connection
- Fast - See [notes on speed](https://github.com/sideshow/apns2/wiki/APNS-HTTP-2-Push-Speed)
- Works with go 1.6 and later
- Supports new Apple Token Based Authentication (JWT)
- Supports new iOS 10 features such as Collapse IDs, Subtitles and Mutable Notifications
- Supports persistent connections to APNs
- Supports VoIP/PushKit notifications (iOS 8 and later)
- Modular & easy to use
- Tested and working in APNs production environment

## Install

- Make sure you have [Go](https://golang.org/doc/install) installed and have set your [GOPATH](https://golang.org/doc/code.html#GOPATH).
- Install apns2:

```sh
go get -u github.com/sideshow/apns2
```

If you are running the test suite you will also need to install testify:
```sh
go get -u github.com/stretchr/testify
```

## Example

```go
package main

import (
  "log"
  "fmt"

  "github.com/sideshow/apns2"
  "github.com/sideshow/apns2/certificate"
)

func main() {

  cert, err := certificate.FromP12File("../cert.p12", "")
  if err != nil {
    log.Fatal("Cert Error:", err)
  }

  notification := &apns2.Notification{}
  notification.DeviceToken = "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"
  notification.Topic = "com.sideshow.Apns2"
  notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below

  client := apns2.NewClient(cert).Production()
  res, err := client.Push(notification)

  if err != nil {
    log.Fatal("Error:", err)
  }

  fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
```

## JWT Token Example

Instead of using a `.p12` or `.pem` certificate as above, you can optionally use
APNs JWT _Provider Authentication Tokens_. First you will need a signing key (`.p8` file), Key ID and Team ID [from Apple](http://help.apple.com/xcode/mac/current/#/dev54d690a66). Once you have these details, you can create a new client:

```go
authKey, err := token.AuthKeyFromFile("../AuthKey_XXX.p8")
if err != nil {
  log.Fatal("token error:", err)
}

token := &token.Token{
  AuthKey: authKey,
  // KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
  KeyID:   "ABC123DEFG",
  // TeamID from developer account (View Account -> Membership)
  TeamID:  "DEF123GHIJ",
}
...

client := apns2.NewTokenClient(token)
res, err := client.Push(notification)
```

- You can use one APNs signing key to authenticate tokens for multiple apps.
- A signing key works for both the development and production environments.
- A signing key doesn’t expire but can be revoked.

## Notification

At a minimum, a _Notification_ needs a _DeviceToken_, a _Topic_ and a _Payload_.

```go
notification := &apns2.Notification{
  DeviceToken: "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7",
  Topic: "com.sideshow.Apns2",
  Payload: []byte(`{"aps":{"alert":"Hello!"}}`),
}
```

You can also set an optional _ApnsID_, _Expiration_ or _Priority_.

```go
notification.ApnsID =  "40636A2C-C093-493E-936A-2A4333C06DEA"
notification.Expiration = time.Now()
notification.Priority = apns2.PriorityLow
```

## Payload

You can use raw bytes for the `notification.Payload` as above, or you can use the payload builder package which makes it easy to construct APNs payloads.

```go
// {"aps":{"alert":"hello","badge":1},"key":"val"}

payload := payload.NewPayload().Alert("hello").Badge(1).Custom("key", "val")

notification.Payload = payload
client.Push(notification)
```

Refer to the [payload](https://godoc.org/github.com/sideshow/apns2/payload) docs for more info.

## Response, Error handling

APNS/2 draws the distinction between a valid response from Apple indicating whether or not the _Notification_ was sent or not, and an unrecoverable or unexpected _Error_;

- An `Error` is returned if a non-recoverable error occurs, i.e. if there is a problem with the underlying _http.Client_ connection or _Certificate_, the payload was not sent, or a valid _Response_ was not received.
- A `Response` is returned if the payload was successfully sent to Apple and a documented response was received. This struct will contain more information about whether or not the push notification succeeded, its _apns-id_ and if applicable, more information around why it did not succeed.

To check if a `Notification` was successfully sent;

```go
res, err := client.Push(notification)
if err != nil {
  log.Println("There was an error", err)
  return
}

if res.Sent() {
  log.Println("Sent:", res.ApnsID)
} else {
  fmt.Printf("Not Sent: %v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
```

## Context & Timeouts

For better control over request cancellations and timeouts APNS/2 supports
contexts. Using a context can be helpful if you want to cancel all pushes when
the parent process is cancelled, or need finer grained control over individual
push timeouts. See the [Google post](https://blog.golang.org/context) for more
information on contexts.

```go
ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
res, err := client.PushWithContext(ctx, notification)
defer cancel()
```

## Speed & Performance

Also see the wiki page on [APNS HTTP 2 Push Speed](https://github.com/sideshow/apns2/wiki/APNS-HTTP-2-Push-Speed).

For best performance, you should hold on to an `apns2.Client` instance and not re-create it every push. The underlying TLS connection itself can take a few seconds to connect and negotiate, so if you are setting up an `apns2.Client` and tearing it down every push, then this will greatly affect performance. (Apple suggest keeping the connection open all the time).

You should also limit the amount of `apns2.Client` instances. The underlying transport has a http connection pool itself, so a single client instance will be enough for most users (One instance can potentially do 4,000+ pushes per second). If you need more than this then one instance per CPU core is a good starting point.

Speed is greatly affected by the location of your server and the quality of your network connection. If you're just testing locally, behind a proxy or if your server is outside USA then you're not going to get great performance. With a good server located in AWS, you should be able to get [decent throughput](https://github.com/sideshow/apns2/wiki/APNS-HTTP-2-Push-Speed).

## Command line tool

APNS/2 has a command line tool that can be installed with `go get github.com/sideshow/apns2/apns2`. Usage:

```
apns2 --help
usage: apns2 --certificate-path=CERTIFICATE-PATH --topic=TOPIC [<flags>]

Listens to STDIN to send notifications and writes APNS response code and reason to STDOUT.

The expected format is: <DeviceToken> <APNS Payload>
Example: aff0c63d9eaa63ad161bafee732d5bc2c31f66d552054718ff19ce314371e5d0 {"aps": {"alert": "hi"}}
Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -c, --certificate-path=CERTIFICATE-PATH
                           Path to certificate file.
  -t, --topic=TOPIC        The topic of the remote notification, which is typically the bundle ID for your app
  -m, --mode="production"  APNS server to send notifications to. `production` or `development`. Defaults to `production`
      --version            Show application version.
```

## License

The MIT License (MIT)

Copyright (c) 2016 Adam Jones

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
