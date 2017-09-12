package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/axiomzen/gorush/config"
	"github.com/axiomzen/gorush/gorush"
	"github.com/axiomzen/gorush/rpc"

	"golang.org/x/sync/errgroup"
)

func main() {
	opts := config.ConfYaml{}

	var (
		showVersion bool
		configFile  string
		topic       string
		message     string
		token       string
		proxy       string
		title       string
	)

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.StringVar(&opts.Core.PID.Path, "pid", "", "PID file path.")
	flag.StringVar(&opts.Ios.KeyPath, "i", "", "iOS certificate key file path")
	flag.StringVar(&opts.Ios.KeyPath, "key", "", "iOS certificate key file path")
	flag.StringVar(&opts.Ios.Password, "P", "", "iOS certificate password for gorush")
	flag.StringVar(&opts.Ios.Password, "password", "", "iOS certificate password for gorush")
	flag.StringVar(&opts.Android.APIKey, "k", "", "Android api key configuration for gorush")
	flag.StringVar(&opts.Android.APIKey, "apikey", "", "Android api key configuration for gorush")
	flag.StringVar(&opts.Core.Address, "A", "", "address to bind")
	flag.StringVar(&opts.Core.Address, "address", "", "address to bind")
	flag.StringVar(&opts.Core.Port, "p", "", "port number for gorush")
	flag.StringVar(&opts.Core.Port, "port", "", "port number for gorush")
	flag.StringVar(&token, "t", "", "token string")
	flag.StringVar(&token, "token", "", "token string")
	flag.StringVar(&opts.Stat.Engine, "e", "", "store engine")
	flag.StringVar(&opts.Stat.Engine, "engine", "", "store engine")
	flag.StringVar(&opts.Stat.Redis.Addr, "redis-addr", "", "redis addr")
	flag.StringVar(&message, "m", "", "notification message")
	flag.StringVar(&message, "message", "", "notification message")
	flag.StringVar(&title, "title", "", "notification title")
	flag.BoolVar(&opts.Android.Enabled, "android", false, "send android notification")
	flag.BoolVar(&opts.Ios.Enabled, "ios", false, "send ios notification")
	flag.BoolVar(&opts.Ios.Production, "production", false, "production mode in iOS")
	flag.StringVar(&topic, "topic", "", "apns topic in iOS")
	flag.StringVar(&proxy, "proxy", "", "http proxy url")

	flag.Usage = usage
	flag.Parse()

	gorush.SetVersion(Version)

	if len(os.Args) < 2 {
		usage()
	}

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

	if opts.Ios.KeyPath != "" {
		gorush.PushConf.Ios.KeyPath = opts.Ios.KeyPath
	}

	if opts.Ios.Password != "" {
		gorush.PushConf.Ios.Password = opts.Ios.Password
	}

	if opts.Android.APIKey != "" {
		gorush.PushConf.Android.APIKey = opts.Android.APIKey
	}

	if opts.Stat.Engine != "" {
		gorush.PushConf.Stat.Engine = opts.Stat.Engine
	}

	if opts.Stat.Redis.Addr != "" {
		gorush.PushConf.Stat.Redis.Addr = opts.Stat.Redis.Addr
	}

	// overwrite server port and address
	if opts.Core.Port != "" {
		gorush.PushConf.Core.Port = opts.Core.Port
	}
	if opts.Core.Address != "" {
		gorush.PushConf.Core.Address = opts.Core.Address
	}

	if err = gorush.InitLog(); err != nil {
		log.Println(err)
		return
	}

	// set http proxy for GCM
	if proxy != "" {
		err = gorush.SetProxy(proxy)

		if err != nil {
			gorush.LogError.Fatal("Set Proxy error: ", err)
		}
	} else if gorush.PushConf.Core.HTTPProxy != "" {
		err = gorush.SetProxy(gorush.PushConf.Core.HTTPProxy)

		if err != nil {
			gorush.LogError.Fatal("Set Proxy error: ", err)
		}
	}

	// send android notification
	if opts.Android.Enabled {
		gorush.PushConf.Android.Enabled = opts.Android.Enabled
		req := gorush.PushNotification{
			Tokens:   []string{token},
			Platform: gorush.PlatFormAndroid,
			Message:  message,
			Title:    title,
		}

		err := gorush.CheckMessage(req)

		if err != nil {
			gorush.LogError.Fatal(err)
		}

		if err := gorush.InitAppStatus(); err != nil {
			return
		}

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
			Title:    title,
		}

		if topic != "" {
			req.Topic = topic
		}

		err := gorush.CheckMessage(req)

		if err != nil {
			gorush.LogError.Fatal(err)
		}

		if err := gorush.InitAppStatus(); err != nil {
			return
		}

		if err := gorush.InitAPNSClient(); err != nil {
			return
		}
		gorush.PushToIOS(req)

		return
	}

	if err = gorush.CheckPushConf(); err != nil {
		gorush.LogError.Fatal(err)
	}

	if opts.Core.PID.Path != "" {
		gorush.PushConf.Core.PID.Path = opts.Core.PID.Path
		gorush.PushConf.Core.PID.Enabled = true
		gorush.PushConf.Core.PID.Override = true
	}

	if err = createPIDFile(); err != nil {
		gorush.LogError.Fatal(err)
	}

	if err = gorush.InitAppStatus(); err != nil {
		return
	}

	gorush.InitWorkers(gorush.PushConf.Core.WorkerNum, gorush.PushConf.Core.QueueNum)

	var g errgroup.Group

	g.Go(func() error {
		return gorush.InitAPNSClient()
	})

	g.Go(func() error {
		_, err := gorush.InitFCMClient(gorush.PushConf.Android.APIKey)
		return err
	})

	g.Go(func() error {
		// Run httpd server
		return gorush.RunHTTPServer()
	})
	g.Go(func() error {
		// Run gRPC internal server
		return rpc.RunGRPCServer()
	})

	if err = g.Wait(); err != nil {
		gorush.LogError.Fatal(err)
	}
}

// Version control for gorush.
var Version = "No Version Provided"

var usageStr = `
  ________                              .__
 /  _____/   ____ _______  __ __  ______|  |__
/   \  ___  /  _ \\_  __ \|  |  \/  ___/|  |  \
\    \_\  \(  <_> )|  | \/|  |  /\___ \ |   Y  \
 \______  / \____/ |__|   |____//____  >|___|  /
        \/                           \/      \/

Usage: gorush [options]

Server Options:
    -A, --address <address>          Address to bind (default: any)
    -p, --port <port>                Use port for clients (default: 8088)
    -c, --config <file>              Configuration file path
    -m, --message <message>          Notification message
    -t, --token <token>              Notification token
    -e, --engine <engine>            Storage engine (memory, redis ...)
    --title <title>                  Notification title
    --proxy <proxy>                  Proxy URL (only for GCM)
    --pid <pid path>                 Process identifier path
    --redis-addr <redis addr>        Redis addr (default: localhost:6379)
iOS Options:
    -i, --key <file>                 certificate key file path
    -P, --password <password>        certificate key password
    --topic <topic>                  iOS topic
    --ios                            enabled iOS (default: false)
    --production                     iOS production mode (default: false)
Android Options:
    -k, --apikey <api_key>           Android API Key
    --android                        enabled android (default: false)
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func createPIDFile() error {
	if !gorush.PushConf.Core.PID.Enabled {
		return nil
	}

	pidPath := gorush.PushConf.Core.PID.Path
	_, err := os.Stat(pidPath)
	if os.IsNotExist(err) || gorush.PushConf.Core.PID.Override {
		currentPid := os.Getpid()
		if err := os.MkdirAll(filepath.Dir(pidPath), os.ModePerm); err != nil {
			return fmt.Errorf("Can't create PID folder on %v", err)
		}

		file, err := os.Create(pidPath)
		if err != nil {
			return fmt.Errorf("Can't create PID file: %v", err)
		}
		defer file.Close()
		if _, err := file.WriteString(strconv.FormatInt(int64(currentPid), 10)); err != nil {
			return fmt.Errorf("Can'write PID information on %s: %v", pidPath, err)
		}
	} else {
		return fmt.Errorf("%s already exists", pidPath)
	}
	return nil
}
