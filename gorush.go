package main

import (
	"flag"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/gorush"
	"log"
	"os"
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
	opts := config.ConfYaml{}

	var showVersion bool
	var configFile string
	var topic string
	var message string
	var token string

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&configFile, "c", "", "Configuration file.")
	flag.StringVar(&configFile, "config", "", "Configuration file.")
	flag.StringVar(&opts.Ios.PemPath, "i", "", "iOS certificate key file path")
	flag.StringVar(&opts.Ios.Password, "password", "", "iOS certificate password for gorush")
	flag.StringVar(&opts.Android.APIKey, "k", "", "Android api key configuration for gorush")
	flag.StringVar(&opts.Core.Port, "p", "", "port number for gorush")
	flag.StringVar(&token, "t", "", "token string")
	flag.StringVar(&message, "m", "", "notification message")
	flag.BoolVar(&opts.Android.Enabled, "android", false, "send android notification")
	flag.BoolVar(&opts.Ios.Enabled, "ios", false, "send ios notification")
	flag.BoolVar(&opts.Ios.Production, "production", false, "production mode in iOS")
	flag.StringVar(&topic, "topic", "", "apns topic in iOS")

	flag.Parse()

	gorush.SetVersion(Version)

	// Show version and exit
	if showVersion {
		gorush.PrintGoRushVersion()
		os.Exit(0)
	}

	var err error

	// set default parameters.
	gorush.PushConf = config.BuildDefaultPushConf()

	// load user define config.
	if configFile != "" {
		gorush.PushConf, err = config.LoadConfYaml(configFile)

		if err != nil {
			log.Printf("Load yaml config file error: '%v'", err)

			return
		}
	}

	if opts.Ios.PemPath != "" {
		gorush.PushConf.Ios.PemPath = opts.Ios.PemPath
	}

	if opts.Ios.Password != "" {
		gorush.PushConf.Ios.Password = opts.Ios.Password
	}

	if opts.Android.APIKey != "" {
		gorush.PushConf.Android.APIKey = opts.Android.APIKey
	}

	// overwrite server port
	if opts.Core.Port != "" {
		gorush.PushConf.Core.Port = opts.Core.Port
	}

	if err = gorush.InitLog(); err != nil {
		log.Println(err)

		return
	}

	// send android notification
	if opts.Android.Enabled {
		gorush.PushConf.Android.Enabled = opts.Android.Enabled
		req := gorush.PushNotification{
			Tokens:   []string{token},
			Platform: gorush.PlatFormAndroid,
			Message:  message,
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
	if opts.Ios.Enabled {
		if opts.Ios.Production {
			gorush.PushConf.Ios.Production = opts.Ios.Production
		}

		gorush.PushConf.Ios.Enabled = opts.Ios.Enabled
		req := gorush.PushNotification{
			Tokens:   []string{token},
			Platform: gorush.PlatFormIos,
			Message:  message,
		}

		if topic != "" {
			req.Topic = topic
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
