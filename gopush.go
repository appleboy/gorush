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
package main

import (
	"flag"
	"github.com/appleboy/gopush/gopush"
	"log"
)

func main() {
	version := flag.Bool("v", false, "gopush version")
	confPath := flag.String("c", "", "yaml configuration file path for gopush")
	certificateKeyPath := flag.String("i", "", "iOS certificate key file path for gopush")
	apiKey := flag.String("k", "", "Android api key configuration for gopush")
	port := flag.String("p", "", "port number for gopush")
	token := flag.String("t", "", "token string")
	message := flag.String("m", "", "notification message")
	android := flag.Bool("android", false, "send android notification")
	ios := flag.Bool("ios", false, "send ios notification")
	production := flag.Bool("production", false, "production mode in iOS")

	flag.Parse()

	if *version {
		gopush.PrintGoPushVersion()
		return
	}

	var err error

	// set default parameters.
	gopush.PushConf = gopush.BuildDefaultPushConf()

	// load user define config.
	if *confPath != "" {
		gopush.PushConf, err = gopush.LoadConfYaml(*confPath)

		if err != nil {
			log.Printf("Load yaml config file error: '%v'", err)

			return
		}
	}

	if *certificateKeyPath != "" {
		gopush.PushConf.Ios.PemKeyPath = *certificateKeyPath
	}

	if *apiKey != "" {
		gopush.PushConf.Android.ApiKey = *apiKey
	}

	// overwrite server port
	if *port != "" {
		gopush.PushConf.Core.Port = *port
	}

	if err = gopush.InitLog(); err != nil {
		log.Println(err)

		return
	}

	// send android notification
	if *android {
		if len(*token) == 0 {
			gopush.LogError.Fatal("Missing token flag (-t)")
		}

		if len(*message) == 0 {
			gopush.LogError.Fatal("Missing message flag (-m)")
		}

		gopush.PushConf.Android.Enabled = true
		req := gopush.PushNotification{
			Tokens:   []string{*token},
			Platform: gopush.PlatFormAndroid,
			Message:  *message,
		}

		gopush.PushToAndroid(req)

		return
	}

	// send android notification
	if *ios {
		if len(*token) == 0 {
			gopush.LogError.Fatal("Missing token flag (-t)")
		}

		if len(*message) == 0 {
			gopush.LogError.Fatal("Missing message flag (-m)")
		}

		if *production {
			gopush.PushConf.Ios.Production = true
		}

		gopush.PushConf.Ios.Enabled = true
		req := gopush.PushNotification{
			Tokens:   []string{*token},
			Platform: gopush.PlatFormIos,
			Message:  *message,
		}

		gopush.InitAPNSClient()
		gopush.PushToIOS(req)

		return
	}

	if err = gopush.CheckPushConf(); err != nil {
		gopush.LogError.Fatal(err)
	}

	gopush.InitAPNSClient()
	gopush.RunHTTPServer()
}
