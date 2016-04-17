// A push notification server using Gin framework written in Go (Golang).
//
// Details about the gorush project are found in github page:
//
//     https://github.com/appleboy/gorush
//
// Support Google Cloud Message using go-gcm library for Android.
// Support HTTP/2 Apple Push Notification Service using apns2 library.
// Support YAML configuration.
// Support command line to send single Android or iOS notification.
// Support Web API to send push notification.
// Support zero downtime restarts for go servers using endless.
// Support HTTP/2 or HTTP/1.1 protocol.
// Support notification queue and multiple workers.
// Support /api/stat/app show notification success and failure counts.
// Support /api/config show your yml config.
//
// The pre-compiled binaries can be downloaded from release page.
//
// Send Android notification
//
//   $ gorush -android -m="your message" -k="API Key" -t="Device token"
//
// Send iOS notification
//
//   $ gorush -ios -m="your message" -i="API Key" -t="Device token"
//
// The default endpoint is APNs development. Please add -production flag for APNs production push endpoint.
//
//   $ gorush -ios -m="your message" -i="API Key" -t="Device token" -production
//
// Run gorush web server
//
//   $ gorush -c config.yml
//
// Get go status of api server using httpie tool:
//
//   $ http -v --verify=no --json GET https://localhost:8088/api/stat/go
//
// Gorush support the following API.
//
// GET /api/stat/go Golang cpu, memory, gc, etc information. Thanks for golang-stats-api-handler.
// GET /api/stat/app show notification success and failure counts.
// GET /api/config show server yml config file.
// POST /api/push push ios and android notifications.
//
// Simple send iOS notification example, the platform value is 1:
//
//   {
//     "notifications": [
//       {
//         "tokens": ["token_a", "token_b"],
//         "platform": 1,
//         "message": "Hello World iOS!"
//       }
//     ]
//   }
//
// Simple send Android notification example, the platform value is 2:
//
//   {
//     "notifications": [
//       {
//         "tokens": ["token_a", "token_b"],
//         "platform": 2,
//         "message": "Hello World Android!"
//       }
//     ]
//   }
//
// For more details, see the documentation and example.
//
package main
