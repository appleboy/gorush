package main

import (
	"flag"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/gorush"
	"log"
)

func checkInput(token, message string) {
	if len(token) == 0 {
		gorush.LogError.Fatal("Missing token flag (-t)")
	}

	if len(message) == 0 {
		gorush.LogError.Fatal("Missing message flag (-m)")
	}
}

var Version = "No Version Provided"

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
	topic := flag.String("topic", "", "apns topic in iOS")

	flag.Parse()

	gorush.SetVersion(Version)

	if *version {
		gorush.PrintGoRushVersion()
		return
	}

	var err error

	// set default parameters.
	gorush.PushConf = config.BuildDefaultPushConf()

	// load user define config.
	if *confPath != "" {
		gorush.PushConf, err = config.LoadConfYaml(*confPath)

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
		gorush.PushConf.Android.Enabled = true
		req := gorush.PushNotification{
			Tokens:   []string{*token},
			Platform: gorush.PlatFormAndroid,
			Message:  *message,
		}

		err := gorush.CheckMessage(req)

		if err != nil {
			gorush.LogError.Fatal(err)
		}

		gorush.InitAppStatus()
		gorush.PushToAndroid(req)

		return
	}

	// send android notification
	if *ios {
		if *production {
			gorush.PushConf.Ios.Production = true
		}

		gorush.PushConf.Ios.Enabled = true
		req := gorush.PushNotification{
			Tokens:   []string{*token},
			Platform: gorush.PlatFormIos,
			Message:  *message,
		}

		if *topic != "" {
			req.Topic = *topic
		}

		err := gorush.CheckMessage(req)

		if err != nil {
			gorush.LogError.Fatal(err)
		}

		gorush.InitAppStatus()
		gorush.InitAPNSClient()
		gorush.PushToIOS(req)

		return
	}

	if err = gorush.CheckPushConf(); err != nil {
		gorush.LogError.Fatal(err)
	}

	gorush.InitAppStatus()
	gorush.InitAPNSClient()
	gorush.InitWorkers(gorush.PushConf.Core.WorkerNum, gorush.PushConf.Core.QueueNum)
	gorush.RunHTTPServer()
}
