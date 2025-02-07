package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/notify"
	"github.com/appleboy/gorush/router"
	"github.com/appleboy/gorush/rpc"
	"github.com/appleboy/gorush/status"

	"github.com/appleboy/graceful"
	"github.com/golang-queue/nats"
	"github.com/golang-queue/nsq"
	"github.com/golang-queue/queue"
	qcore "github.com/golang-queue/queue/core"
	redisdb "github.com/golang-queue/redisdb-stream"
)

//nolint:gocyclo
func main() {
	opts := config.ConfYaml{}

	var (
		ping        bool
		showVersion bool
		configFile  string
		topic       string
		message     string
		token       string
		title       string
	)

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "V", false, "Print version information.")
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.StringVar(&opts.Core.PID.Path, "pid", "", "PID file path.")
	flag.StringVar(&opts.Ios.KeyPath, "i", "", "iOS certificate key file path")
	flag.StringVar(&opts.Ios.KeyPath, "key", "", "iOS certificate key file path")
	flag.StringVar(&opts.Ios.KeyID, "key-id", "", "iOS Key ID for P8 token")
	flag.StringVar(&opts.Ios.TeamID, "team-id", "", "iOS Team ID for P8 token")
	flag.StringVar(&opts.Ios.Password, "P", "", "iOS certificate password for gorush")
	flag.StringVar(&opts.Ios.Password, "password", "", "iOS certificate password for gorush")
	flag.StringVar(&opts.Android.KeyPath, "fcm-key", "", "FCM key path configuration for gorush")
	flag.StringVar(&opts.Huawei.AppSecret, "hk", "", "Huawei api key configuration for gorush")
	flag.StringVar(&opts.Huawei.AppSecret, "hmskey", "", "Huawei api key configuration for gorush")
	flag.StringVar(&opts.Huawei.AppID, "hid", "", "HMS app id configuration for gorush")
	flag.StringVar(&opts.Huawei.AppID, "hmsid", "", "HMS app id configuration for gorush")
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
	flag.BoolVar(&opts.Huawei.Enabled, "huawei", false, "send huawei notification")
	flag.BoolVar(&opts.Ios.Enabled, "ios", false, "send ios notification")
	flag.BoolVar(&opts.Ios.Production, "production", false, "production mode in iOS")
	flag.StringVar(&topic, "topic", "", "apns topic in iOS")
	flag.StringVar(&opts.Core.HTTPProxy, "proxy", "", "http proxy url")
	flag.BoolVar(&ping, "ping", false, "ping server")

	flag.Usage = usage
	flag.Parse()

	router.SetVersion(version)
	router.SetCommit(commit)

	// Show version and exit
	if showVersion {
		router.PrintGoRushVersion()
		os.Exit(0)
	}

	// set default parameters.
	cfg, err := config.LoadConf(configFile)
	if err != nil {
		log.Printf("Load yaml config file error: '%v'", err)

		return
	}

	// Initialize push slots for concurrent iOS pushes
	notify.MaxConcurrentIOSPushes = make(chan struct{}, cfg.Ios.MaxConcurrentPushes)

	if opts.Ios.KeyPath != "" {
		cfg.Ios.KeyPath = opts.Ios.KeyPath
	}

	if opts.Ios.KeyID != "" {
		cfg.Ios.KeyID = opts.Ios.KeyID
	}

	if opts.Ios.TeamID != "" {
		cfg.Ios.TeamID = opts.Ios.TeamID
	}

	if opts.Ios.Password != "" {
		cfg.Ios.Password = opts.Ios.Password
	}

	if opts.Android.KeyPath != "" {
		cfg.Android.KeyPath = opts.Android.KeyPath
	}

	if opts.Huawei.AppSecret != "" {
		cfg.Huawei.AppSecret = opts.Huawei.AppSecret
	}

	if opts.Huawei.AppID != "" {
		cfg.Huawei.AppID = opts.Huawei.AppID
	}

	if opts.Stat.Engine != "" {
		cfg.Stat.Engine = opts.Stat.Engine
	}

	if opts.Stat.Redis.Addr != "" {
		cfg.Stat.Redis.Addr = opts.Stat.Redis.Addr
	}

	// overwrite server port and address
	if opts.Core.Port != "" {
		cfg.Core.Port = opts.Core.Port
	}
	if opts.Core.Address != "" {
		cfg.Core.Address = opts.Core.Address
	}

	if err = logx.InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	); err != nil {
		log.Fatalf("can't load log module, error: %v", err)
	}

	if opts.Core.HTTPProxy != "" {
		cfg.Core.HTTPProxy = opts.Core.HTTPProxy
	}

	if cfg.Core.HTTPProxy != "" {
		err = notify.SetProxy(cfg.Core.HTTPProxy)
		if err != nil {
			logx.LogError.Fatalf("Set Proxy error: %v", err)
		}
	}

	g := graceful.NewManager(
		graceful.WithLogger(logx.QueueLogger()),
	)

	if ping {
		if err := pinger(g.ShutdownContext(), cfg); err != nil {
			logx.LogError.Fatal(err)
		}
		return
	}

	// send android notification
	if opts.Android.Enabled {
		cfg.Android.Enabled = opts.Android.Enabled
		req := &notify.PushNotification{
			Platform: core.PlatFormAndroid,
			Message:  message,
			Title:    title,
		}

		// send message to single device
		if token != "" {
			req.To = token
		}

		// send topic message
		if topic != "" {
			req.Topic = topic
		}

		if err := status.InitAppStatus(cfg); err != nil {
			return
		}

		if _, err := notify.PushToAndroid(g.ShutdownContext(), req, cfg); err != nil {
			return
		}

		return
	}

	// send huawei notification
	if opts.Huawei.Enabled {
		cfg.Huawei.Enabled = opts.Huawei.Enabled
		req := &notify.PushNotification{
			Platform: core.PlatFormHuawei,
			Message:  message,
			Title:    title,
		}

		// send message to single device
		if token != "" {
			req.Tokens = []string{token}
		}

		// send topic message
		if topic != "" {
			req.To = topic
		}

		err := notify.CheckMessage(req)
		if err != nil {
			logx.LogError.Fatal(err)
		}

		if err := status.InitAppStatus(cfg); err != nil {
			return
		}

		if _, err := notify.PushToHuawei(g.ShutdownContext(), req, cfg); err != nil {
			return
		}

		return
	}

	// send ios notification
	if opts.Ios.Enabled {
		if opts.Ios.Production {
			cfg.Ios.Production = opts.Ios.Production
		}

		cfg.Ios.Enabled = opts.Ios.Enabled
		req := &notify.PushNotification{
			Platform: core.PlatFormIos,
			Message:  message,
			Title:    title,
		}

		// send message to single device
		if token != "" {
			req.Tokens = []string{token}
		}

		// send topic message
		if topic != "" {
			req.Topic = topic
		}

		err := notify.CheckMessage(req)
		if err != nil {
			logx.LogError.Fatal(err)
		}

		if err := status.InitAppStatus(cfg); err != nil {
			return
		}

		if err := notify.InitAPNSClient(g.ShutdownContext(), cfg); err != nil {
			return
		}

		if _, err := notify.PushToIOS(g.ShutdownContext(), req, cfg); err != nil {
			return
		}

		return
	}

	if err = notify.CheckPushConf(cfg); err != nil {
		logx.LogError.Fatal(err)
	}

	if opts.Core.PID.Path != "" {
		cfg.Core.PID.Path = opts.Core.PID.Path
		cfg.Core.PID.Enabled = true
		cfg.Core.PID.Override = true
	}

	if err = createPIDFile(cfg); err != nil {
		logx.LogError.Fatal(err)
	}

	if err = status.InitAppStatus(cfg); err != nil {
		logx.LogError.Fatal(err)
	}

	var w qcore.Worker
	switch core.Queue(cfg.Queue.Engine) {
	case core.LocalQueue:
		w = queue.NewRing(
			queue.WithQueueSize(int(cfg.Core.QueueNum)),
			queue.WithFn(notify.Run(cfg)),
			queue.WithLogger(logx.QueueLogger()),
		)
	case core.NSQ:
		w = nsq.NewWorker(
			nsq.WithAddr(cfg.Queue.NSQ.Addr),
			nsq.WithTopic(cfg.Queue.NSQ.Topic),
			nsq.WithChannel(cfg.Queue.NSQ.Channel),
			nsq.WithMaxInFlight(int(cfg.Core.WorkerNum)),
			nsq.WithRunFunc(notify.Run(cfg)),
			nsq.WithLogger(logx.QueueLogger()),
		)
	case core.NATS:
		w = nats.NewWorker(
			nats.WithAddr(cfg.Queue.NATS.Addr),
			nats.WithSubj(cfg.Queue.NATS.Subj),
			nats.WithQueue(cfg.Queue.NATS.Queue),
			nats.WithRunFunc(notify.Run(cfg)),
			nats.WithLogger(logx.QueueLogger()),
		)
	case core.Redis:
		opts := []redisdb.Option{
			redisdb.WithAddr(cfg.Queue.Redis.Addr),
			redisdb.WithUsername(cfg.Queue.Redis.Username),
			redisdb.WithPassword(cfg.Queue.Redis.Password),
			redisdb.WithStreamName(cfg.Queue.Redis.StreamName),
			redisdb.WithGroup(cfg.Queue.Redis.Group),
			redisdb.WithConsumer(cfg.Queue.Redis.Consumer),
			redisdb.WithMaxLength(cfg.Core.QueueNum),
			redisdb.WithRunFunc(notify.Run(cfg)),
			redisdb.WithLogger(logx.QueueLogger()),
		}
		if cfg.Queue.Redis.WithTLS {
			opts = append(opts, redisdb.WithTLS())
		}
		w = redisdb.NewWorker(
			opts...,
		)
	default:
		logx.LogError.Fatalf("we don't support queue engine: %s", cfg.Queue.Engine)
	}

	q := queue.NewPool(
		cfg.Core.WorkerNum,
		queue.WithWorker(w),
		queue.WithLogger(logx.QueueLogger()),
	)

	g.AddShutdownJob(func() error {
		// logx.LogAccess.Info("close the queue system, current queue usage: ", q.Usage())
		// stop queue system and wait job completed
		q.Release()
		// close the connection with storage
		logx.LogAccess.Info("close the storage connection: ", cfg.Stat.Engine)
		if err := status.StatStorage.Close(); err != nil {
			logx.LogError.Fatal("can't close the storage connection: ", err.Error())
		}
		return nil
	})

	if cfg.Ios.Enabled {
		if err = notify.InitAPNSClient(g.ShutdownContext(), cfg); err != nil {
			logx.LogError.Fatal(err)
		}
	}

	if cfg.Android.Enabled {
		if _, err = notify.InitFCMClient(g.ShutdownContext(), cfg); err != nil {
			logx.LogError.Fatal(err)
		}
	}

	if cfg.Huawei.Enabled {
		if _, err = notify.InitHMSClient(cfg, cfg.Huawei.AppSecret, cfg.Huawei.AppID); err != nil {
			logx.LogError.Fatal(err)
		}
	}

	g.AddRunningJob(func(ctx context.Context) error {
		return router.RunHTTPServer(ctx, cfg, q)
	})

	g.AddRunningJob(func(ctx context.Context) error {
		return rpc.RunGRPCServer(ctx, cfg)
	})

	<-g.Done()
}

// Version control for notify.
var (
	version = "No Version Provided"
	commit  = "No Commit Provided"
)

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
    --proxy <proxy>                  Proxy URL
    --pid <pid path>                 Process identifier path
    --redis-addr <redis addr>        Redis addr (default: localhost:6379)
    --ping                           healthy check command for container
iOS Options:
    -i, --key <file>                 certificate key file path
    -P, --password <password>        certificate key password
    --ios                            enabled iOS (default: false)
    --production                     iOS production mode (default: false)
Android Options:
    --fcm-key <fcm_key_path>         FCM Credentials Key Path
    --android                        enabled android (default: false)
Huawei Options:
    -hk, --hmskey <hms_key>          HMS App Secret
    -hid, --hmsid <hms_id>           HMS App ID
    --huawei                         enabled huawei (default: false)
Common Options:
    --topic <topic>                  iOS, Android or Huawei topic message
    -h, --help                       Show this message
    -V, --version                    Show version
`

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
}

// handles pinging the endpoint and returns an error if the
// agent is in an unhealthy state.
func pinger(ctx context.Context, cfg *config.ConfYaml) error {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	req, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://localhost:"+cfg.Core.Port+cfg.API.HealthURI,
		nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status code")
	}
	return nil
}

func createPIDFile(cfg *config.ConfYaml) error {
	if !cfg.Core.PID.Enabled {
		return nil
	}

	pidPath := cfg.Core.PID.Path
	_, err := os.Stat(pidPath)
	if os.IsNotExist(err) || cfg.Core.PID.Override {
		currentPid := os.Getpid()
		if err := os.MkdirAll(filepath.Dir(pidPath), os.ModePerm); err != nil {
			return fmt.Errorf("can't create PID folder on %v", err)
		}

		file, err := os.Create(pidPath)
		if err != nil {
			return fmt.Errorf("can't create PID file: %v", err)
		}
		defer file.Close()
		if _, err := file.WriteString(strconv.FormatInt(int64(currentPid), 10)); err != nil {
			return fmt.Errorf("can't write PID information on %s: %v", pidPath, err)
		}
	} else {
		return fmt.Errorf("%s already exists", pidPath)
	}
	return nil
}
