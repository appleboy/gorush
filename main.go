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

	"github.com/appleboy/gorush/app"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/notify"
	"github.com/appleboy/gorush/router"
	"github.com/appleboy/gorush/rpc"
	"github.com/appleboy/gorush/status"

	"github.com/appleboy/graceful"
)

func main() {
	// Parse CLI flags
	opts := app.NewOptions()
	opts.BindFlags()
	flag.Usage = usage
	flag.Parse()

	router.SetVersion(version)
	router.SetCommit(commit)

	// Show version and exit
	if opts.ShowVersion {
		router.PrintGoRushVersion()
		os.Exit(0)
	}

	// Load and merge configuration
	cfg, err := app.ValidateAndMerge(opts)
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Initialize push slots for concurrent iOS pushes
	notify.MaxConcurrentIOSPushes = make(chan struct{}, cfg.Ios.MaxConcurrentPushes)

	if err = logx.InitLog(
		cfg.Log.AccessLevel,
		cfg.Log.AccessLog,
		cfg.Log.ErrorLevel,
		cfg.Log.ErrorLog,
	); err != nil {
		log.Fatalf("can't load log module, error: %v", err)
	}

	if cfg.Core.HTTPProxy != "" {
		if err = notify.SetProxy(cfg.Core.HTTPProxy); err != nil {
			logx.LogError.Fatalf("Set Proxy error: %v", err)
		}
	}

	g := graceful.NewManager(
		graceful.WithLogger(logx.QueueLogger()),
	)

	if opts.Ping {
		if err := pinger(g.ShutdownContext(), cfg); err != nil {
			logx.LogError.Fatal(err)
		}
		return
	}

	// Handle CLI notification mode
	if opts.IsCLIMode() {
		if err := handleCLINotification(g.ShutdownContext(), cfg, opts); err != nil {
			logx.LogError.Fatalf("Failed to send notification: %v", err)
		}
		return
	}

	if err = notify.CheckPushConf(cfg); err != nil {
		logx.LogError.Fatal(err)
	}

	if err = createPIDFile(cfg); err != nil {
		logx.LogError.Fatal(err)
	}

	if err = status.InitAppStatus(cfg); err != nil {
		logx.LogError.Fatal(err)
	}

	w, err := app.NewQueueWorker(cfg)
	if err != nil {
		logx.LogError.Fatal(err)
	}

	q := app.NewQueuePool(cfg, w)

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

// handleCLINotification handles sending notifications in CLI mode.
func handleCLINotification(ctx context.Context, cfg *config.ConfYaml, opts *app.Options) error {
	sendOpts := opts.CLISendOptions()

	if opts.Conf.Android.Enabled {
		return app.SendAndroidNotification(ctx, cfg, sendOpts)
	}

	if opts.Conf.Huawei.Enabled {
		return app.SendHuaweiNotification(ctx, cfg, sendOpts)
	}

	if opts.Conf.Ios.Enabled {
		return app.SendIOSNotification(ctx, cfg, sendOpts)
	}

	return nil
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

	// Additional validation before creating PID file
	if err := config.ValidatePIDPath(pidPath); err != nil {
		return err
	}
	_, err := os.Stat(pidPath)
	if os.IsNotExist(err) || cfg.Core.PID.Override {
		currentPid := os.Getpid()
		if err := os.MkdirAll(filepath.Dir(pidPath), os.ModePerm); err != nil {
			return fmt.Errorf("can't create PID folder: %w", err)
		}

		file, err := os.Create(pidPath)
		if err != nil {
			return fmt.Errorf("can't create PID file: %w", err)
		}
		defer file.Close()
		if _, err := file.WriteString(strconv.FormatInt(int64(currentPid), 10)); err != nil {
			return fmt.Errorf("can't write PID information on %s: %w", pidPath, err)
		}
	} else {
		return fmt.Errorf("%s already exists", pidPath)
	}
	return nil
}
