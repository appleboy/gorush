package main

import (
	"flag"
	"github.com/appleboy/gorush/gorush"
	"log"
)

func main() {
	version := flag.Bool("v", false, "gorush version")
	confPath := flag.String("c", "", "yaml configuration file path for gorush")
	certificateKeyPath := flag.String("i", "", "iOS certificate key file path for gorush")
	apiKey := flag.String("k", "", "Android api key configuration for gorush")
	port := flag.String("p", "", "port number for gorush")
	token := flag.String("t", "", "token string")
	message := flag.String("m", "", "notification message")
	android := flag.Bool("android", false, "send android notification")
	ios := flag.Bool("ios", false, "send ios notification")
	production := flag.Bool("production", false, "production mode in iOS")

	flag.Parse()

	if *version {
		gorush.PrintGoRushVersion()
		return
	}

	var err error

	// set default parameters.
	gorush.PushConf = gorush.BuildDefaultPushConf()

	// load user define config.
	if *confPath != "" {
		gorush.PushConf, err = gorush.LoadConfYaml(*confPath)

		if err != nil {
			log.Printf("Load yaml config file error: '%v'", err)

			return
		}
	}

	if *certificateKeyPath != "" {
		gorush.PushConf.Ios.PemKeyPath = *certificateKeyPath
	}

	if *apiKey != "" {
		gorush.PushConf.Android.APIKey = *apiKey
	}

	// overwrite server port
	if *port != "" {
		gorush.PushConf.Core.Port = *port
	}

	if err = gorush.InitLog(); err != nil {
		log.Println(err)

		return
	}

	// send android notification
	if *android {
		if len(*token) == 0 {
			gorush.LogError.Fatal("Missing token flag (-t)")
		}

		if len(*message) == 0 {
			gorush.LogError.Fatal("Missing message flag (-m)")
		}

		gorush.PushConf.Android.Enabled = true
		req := gorush.PushNotification{
			Tokens:   []string{*token},
			Platform: gorush.PlatFormAndroid,
			Message:  *message,
		}

		gorush.PushToAndroid(req)

		return
	}

	// send android notification
	if *ios {
		if len(*token) == 0 {
			gorush.LogError.Fatal("Missing token flag (-t)")
		}

		if len(*message) == 0 {
			gorush.LogError.Fatal("Missing message flag (-m)")
		}

		if *production {
			gorush.PushConf.Ios.Production = true
		}

		gorush.PushConf.Ios.Enabled = true
		req := gorush.PushNotification{
			Tokens:   []string{*token},
			Platform: gorush.PlatFormIos,
			Message:  *message,
		}

		gorush.InitAPNSClient()
		gorush.PushToIOS(req)

		return
	}

	if err = gorush.CheckPushConf(); err != nil {
		gorush.LogError.Fatal(err)
	}

	gorush.InitAPNSClient()
	gorush.RunHTTPServer()
}
