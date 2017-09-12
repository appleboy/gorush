// A push notification server using Gin framework written in Go (Golang).
//
// Details about the gorush project are found in github page:
//
//     https://github.com/axiomzen/gorush
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
