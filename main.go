package main

import (
	"flag"
	"github.com/appleboy/gopush/gopush"
	"github.com/sideshow/apns2/certificate"
	apns "github.com/sideshow/apns2"
	"log"
)

func main() {
	version := flag.Bool("v", false, "gopush version")
	confPath := flag.String("c", "", "yaml configuration file path for gopush")
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

	if gopush.PushConf.Ios.Enabled {
		gopush.CertificatePemIos, err = certificate.FromPemFile(gopush.PushConf.Ios.PemKeyPath, "")

		if err != nil {
			log.Println("Cert Error:", err)

			return
		}

		if gopush.PushConf.Ios.Production {
			gopush.ApnsClient = apns.NewClient(gopush.CertificatePemIos).Production()
		} else {
			gopush.ApnsClient = apns.NewClient(gopush.CertificatePemIos).Development()
		}
	}

	// overwrite server port
	if *port != "" {
		gopush.PushConf.Core.Port = *port
	}

	gopush.RunHTTPServer()
}
