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
			log.Printf("Unable to load yaml config file: '%v'", err)

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

	gopush.InitLog()

	if err = gopush.CheckPushConf(); err != nil {
		log.Println(err)
		gopush.LogError.Fatal(err)

		return
	}

	gopush.InitAPNSClient()
	gopush.RunHTTPServer()
}
