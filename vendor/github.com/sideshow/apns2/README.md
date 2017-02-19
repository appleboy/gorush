# APNS/2

APNS/2 is a go package designed for simple, flexible and fast Apple Push Notifications on iOS, OSX and Safari using the new HTTP/2 Push provider API.

[![Build Status](https://travis-ci.org/sideshow/apns2.svg?branch=master)](https://travis-ci.org/sideshow/apns2)  [![Coverage Status](https://coveralls.io/repos/sideshow/apns2/badge.svg?branch=master&service=github)](https://coveralls.io/github/sideshow/apns2?branch=master)  [![GoDoc](https://godoc.org/github.com/sideshow/apns2?status.svg)](https://godoc.org/github.com/sideshow/apns2)

## Features

- Uses new Apple APNs HTTP/2 connection
- Works with go 1.6 and later
- Supports new iOS 10 features such as Collapse IDs, Subtitles and Mutable Notifications
- Supports persistent connections to APNs
- Supports VoIP/PushKit notifications (iOS 8 and later)
- Fast, modular & easy to use
- Tested and working in APNs production environment

## Install

- Make sure you have [Go](https://golang.org/doc/install) installed and have set your [GOPATH](https://golang.org/doc/code.html#GOPATH).
- Download and install the dependencies:

  ```sh
go get -u golang.org/x/net/http2
go get -u golang.org/x/crypto/pkcs12
  ```

- Install apns2:

  ```sh
go get -u github.com/sideshow/apns2
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

## Notification

At a minimum, a _Notification_ needs a _DeviceToken_, a _Topic_ and a _Payload_.

```go
notification := &Notification{
  DeviceToken: "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7",
  Topic: "com.sideshow.Apns2",
  Payload: []byte(`{"aps":{"alert":"Hello!"}}`),
}
```

You can also set an optional _ApnsID_, _Expiration_ or _Priority_.

```go
notification.ApnsID =  "40636A2C-C093-493E-936A-2A4333C06DEA"
notification.Expiration = time.Now()
notification.Priority = apns.PriorityLow
```

## Payload

You can use raw bytes for the `notification.Payload` as above, or you can use the payload builder package which makes it easy to construct APNs payloads.

```go
// {"aps":{"alert":"hello","badge":1},"key":"val"}

payload := NewPayload().Alert("hello").Badge(1).Custom("key", "val")

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
