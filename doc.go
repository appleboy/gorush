// A push notification server using Gin framework written in Go (Golang).
//
// Details about the gopush project are found in github page:
//
//     https://github.com/appleboy/gopush
//
// Support Google Cloud Message using go-gcm library for Android.
// Support HTTP/2 Apple Push Notification Service using apns2 library.
// Support YAML configuration.
// Support command line to send single Android or iOS notification.
// Support Web API to send push notification.
// Support zero downtime restarts for go servers using endless.
// Support HTTP/2 or HTTP/1.1 protocol.
//
// The pre-compiled binaries can be downloaded from release page.
//
// Send Android notification
//
//   $ gopush -android -m="your message" -k="API Key" -t="Device token"
//
// Send iOS notification
//
//   $ gopush -ios -m="your message" -i="API Key" -t="Device token"
//
// The default endpoint is APNs development. Please add -production flag for APNs production push endpoint.
//
//   $ gopush -ios -m="your message" -i="API Key" -t="Device token" -production
//
// For more details, see the documentation and example.
//
