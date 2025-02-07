package config

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
core:
  enabled: true # enable httpd server
  address: "" # ip address to bind (default: any)
  shutdown_timeout: 30 # default is 30 second
  port: "8088" # ignore this port number if auto_tls is enabled (listen 443).
  worker_num: 0 # default worker number is runtime.NumCPU()
  queue_num: 0 # default queue number is 8192
  max_notification: 100
  # set true if you need get error message from fail push notification in API response.
  # It only works when the queue engine is local.
  sync: false
  # set webhook url if you need get error message asynchronously from fail push notification in API response.
  feedback_hook_url: ""
  feedback_timeout: 10 # default is 10 second
  feedback_header:
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"
  cert_base64: ""
  key_base64: ""
  http_proxy: ""
  pid:
    enabled: false
    path: "gorush.pid"
    override: true
  auto_tls:
    enabled: false # Automatically install TLS certificates from Let's Encrypt.
    folder: ".cache" # folder for storing TLS certificates
    host: "" # which domains the Let's Encrypt will attempt

grpc:
  enabled: false # enable gRPC server
  port: 9000

api:
  push_uri: "/api/push"
  stat_go_uri: "/api/stat/go"
  stat_app_uri: "/api/stat/app"
  config_uri: "/api/config"
  sys_stat_uri: "/sys/stats"
  metric_uri: "/metrics"
  health_uri: "/healthz"

android:
  enabled: true
  key_path: "" # path to fcm key file
  credential: "" # fcm credential data
  max_retry: 0 # resend fail notification, default value zero is disabled

huawei:
  enabled: false
  appsecret: "YOUR_APP_SECRET"
  appid: "YOUR_APP_ID"
  max_retry: 0 # resend fail notification, default value zero is disabled

queue:
  engine: "local" # support "local", "nsq", "nats" and "redis" default value is "local"
  nsq:
    addr: 127.0.0.1:4150
    topic: gorush
    channel: gorush
  nats:
    addr: 127.0.0.1:4222
    subj: gorush
    queue: gorush
  redis:
    addr: 127.0.0.1:6379
    group: gorush
    consumer: gorush
    stream_name: gorush
    with_tls: false
    username: ""
    password: ""

ios:
  enabled: false
  key_path: ""
  key_base64: "" # load iOS key from base64 input
  key_type: "pem" # could be pem, p12 or p8 type
  password: "" # certificate password, default as empty string.
  production: false
  max_concurrent_pushes: 100 # just for push ios notification
  max_retry: 0 # resend fail notification, default value zero is disabled
  key_id: "" # KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
  team_id: "" # TeamID from developer account (View Account -> Membership)

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console, or define log path like "log/error_log"
  error_level: "error"
  hide_token: true
  hide_messages: false

stat:
  engine: "memory" # support memory, redis, boltdb, buntdb or leveldb
  redis:
    cluster: false
    addr: "localhost:6379" # if cluster is true, you may set this to "localhost:6379,localhost:6380,localhost:6381"
    username: ""
    password: ""
    db: 0
  boltdb:
    path: "bolt.db"
    bucket: "gorush"
  buntdb:
    path: "bunt.db"
  leveldb:
    path: "level.db"
  badgerdb:
    path: "badger.db"
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core    SectionCore    `yaml:"core"`
	API     SectionAPI     `yaml:"api"`
	Android SectionAndroid `yaml:"android"`
	Huawei  SectionHuawei  `yaml:"huawei"`
	Ios     SectionIos     `yaml:"ios"`
	Queue   SectionQueue   `yaml:"queue"`
	Log     SectionLog     `yaml:"log"`
	Stat    SectionStat    `yaml:"stat"`
	GRPC    SectionGRPC    `yaml:"grpc"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Enabled         bool           `yaml:"enabled"`
	Address         string         `yaml:"address"`
	ShutdownTimeout int64          `yaml:"shutdown_timeout"`
	Port            string         `yaml:"port"`
	MaxNotification int64          `yaml:"max_notification"`
	WorkerNum       int64          `yaml:"worker_num"`
	QueueNum        int64          `yaml:"queue_num"`
	Mode            string         `yaml:"mode"`
	Sync            bool           `yaml:"sync"`
	SSL             bool           `yaml:"ssl"`
	CertPath        string         `yaml:"cert_path"`
	KeyPath         string         `yaml:"key_path"`
	CertBase64      string         `yaml:"cert_base64"`
	KeyBase64       string         `yaml:"key_base64"`
	HTTPProxy       string         `yaml:"http_proxy"`
	PID             SectionPID     `yaml:"pid"`
	AutoTLS         SectionAutoTLS `yaml:"auto_tls"`

	FeedbackURL     string   `yaml:"feedback_hook_url"`
	FeedbackTimeout int64    `yaml:"feedback_timeout"`
	FeedbackHeader  []string `yaml:"feedback_header"`
}

// SectionAutoTLS support Let's Encrypt setting.
type SectionAutoTLS struct {
	Enabled bool   `yaml:"enabled"`
	Folder  string `yaml:"folder"`
	Host    string `yaml:"host"`
}

// SectionAPI is sub section of config.
type SectionAPI struct {
	PushURI    string `yaml:"push_uri"`
	StatGoURI  string `yaml:"stat_go_uri"`
	StatAppURI string `yaml:"stat_app_uri"`
	ConfigURI  string `yaml:"config_uri"`
	SysStatURI string `yaml:"sys_stat_uri"`
	MetricURI  string `yaml:"metric_uri"`
	HealthURI  string `yaml:"health_uri"`
}

// SectionAndroid is sub section of config.
type SectionAndroid struct {
	Enabled    bool   `yaml:"enabled"`
	KeyPath    string `yaml:"key_path"`
	Credential string `yaml:"credential"`
	MaxRetry   int    `yaml:"max_retry"`
}

// SectionHuawei is sub section of config.
type SectionHuawei struct {
	Enabled   bool   `yaml:"enabled"`
	AppSecret string `yaml:"appsecret"`
	AppID     string `yaml:"appid"`
	MaxRetry  int    `yaml:"max_retry"`
}

// SectionIos is sub section of config.
type SectionIos struct {
	Enabled             bool   `yaml:"enabled"`
	KeyPath             string `yaml:"key_path"`
	KeyBase64           string `yaml:"key_base64"`
	KeyType             string `yaml:"key_type"`
	Password            string `yaml:"password"`
	Production          bool   `yaml:"production"`
	MaxConcurrentPushes uint   `yaml:"max_concurrent_pushes"`
	MaxRetry            int    `yaml:"max_retry"`
	KeyID               string `yaml:"key_id"`
	TeamID              string `yaml:"team_id"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Format       string `yaml:"format"`
	AccessLog    string `yaml:"access_log"`
	AccessLevel  string `yaml:"access_level"`
	ErrorLog     string `yaml:"error_log"`
	ErrorLevel   string `yaml:"error_level"`
	HideToken    bool   `yaml:"hide_token"`
	HideMessages bool   `yaml:"hide_messages"`
}

// SectionStat is sub section of config.
type SectionStat struct {
	Engine   string          `yaml:"engine"`
	Redis    SectionRedis    `yaml:"redis"`
	BoltDB   SectionBoltDB   `yaml:"boltdb"`
	BuntDB   SectionBuntDB   `yaml:"buntdb"`
	LevelDB  SectionLevelDB  `yaml:"leveldb"`
	BadgerDB SectionBadgerDB `yaml:"badgerdb"`
}

// SectionQueue is sub section of config.
type SectionQueue struct {
	Engine string            `yaml:"engine"`
	NSQ    SectionNSQ        `yaml:"nsq"`
	NATS   SectionNATS       `yaml:"nats"`
	Redis  SectionRedisQueue `yaml:"redis"`
}

// SectionNSQ is sub section of config.
type SectionNSQ struct {
	Addr    string `yaml:"addr"`
	Topic   string `yaml:"topic"`
	Channel string `yaml:"channel"`
}

// SectionNATS is sub section of config.
type SectionNATS struct {
	Addr  string `yaml:"addr"`
	Subj  string `yaml:"subj"`
	Queue string `yaml:"queue"`
}

// SectionRedisQueue is sub section of config.
type SectionRedisQueue struct {
	Addr       string `yaml:"addr"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	StreamName string `yaml:"stream_name"`
	Group      string `yaml:"group"`
	Consumer   string `yaml:"consumer"`
	WithTLS    bool   `yaml:"with_tls"`
}

// SectionRedis is sub section of config.
type SectionRedis struct {
	Cluster  bool   `yaml:"cluster"`
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// SectionBoltDB is sub section of config.
type SectionBoltDB struct {
	Path   string `yaml:"path"`
	Bucket string `yaml:"bucket"`
}

// SectionBuntDB is sub section of config.
type SectionBuntDB struct {
	Path string `yaml:"path"`
}

// SectionLevelDB is sub section of config.
type SectionLevelDB struct {
	Path string `yaml:"path"`
}

// SectionBadgerDB is sub section of config.
type SectionBadgerDB struct {
	Path string `yaml:"path"`
}

// SectionPID is sub section of config.
type SectionPID struct {
	Enabled  bool   `yaml:"enabled"`
	Path     string `yaml:"path"`
	Override bool   `yaml:"override"`
}

// SectionGRPC is sub section of config.
type SectionGRPC struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

func setDefault() {
	viper.SetDefault("ios.max_concurrent_pushes", uint(100))
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath ...string) (*ConfYaml, error) {
	conf := &ConfYaml{}

	// load default values
	setDefault()

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("gorush") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if len(confPath) > 0 && confPath[0] != "" {
		content, err := os.ReadFile(confPath[0])
		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".gorush" (without extension).
		viper.AddConfigPath("/etc/gorush/")
		viper.AddConfigPath("$HOME/.gorush")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
			// load default config
			return conf, err
		}
	}

	// Core
	conf.Core.Address = viper.GetString("core.address")
	conf.Core.Port = viper.GetString("core.port")
	conf.Core.ShutdownTimeout = int64(viper.GetInt("core.shutdown_timeout"))
	conf.Core.Enabled = viper.GetBool("core.enabled")
	conf.Core.WorkerNum = int64(viper.GetInt("core.worker_num"))
	conf.Core.QueueNum = int64(viper.GetInt("core.queue_num"))
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.Sync = viper.GetBool("core.sync")
	conf.Core.FeedbackURL = viper.GetString("core.feedback_hook_url")
	conf.Core.FeedbackTimeout = int64(viper.GetInt("core.feedback_timeout"))
	conf.Core.FeedbackHeader = viper.GetStringSlice("core.feedback_header")
	conf.Core.SSL = viper.GetBool("core.ssl")
	conf.Core.CertPath = viper.GetString("core.cert_path")
	conf.Core.KeyPath = viper.GetString("core.key_path")
	conf.Core.CertBase64 = viper.GetString("core.cert_base64")
	conf.Core.KeyBase64 = viper.GetString("core.key_base64")
	conf.Core.MaxNotification = int64(viper.GetInt("core.max_notification"))
	conf.Core.HTTPProxy = viper.GetString("core.http_proxy")
	conf.Core.PID.Enabled = viper.GetBool("core.pid.enabled")
	conf.Core.PID.Path = viper.GetString("core.pid.path")
	conf.Core.PID.Override = viper.GetBool("core.pid.override")
	conf.Core.AutoTLS.Enabled = viper.GetBool("core.auto_tls.enabled")
	conf.Core.AutoTLS.Folder = viper.GetString("core.auto_tls.folder")
	conf.Core.AutoTLS.Host = viper.GetString("core.auto_tls.host")

	// Api
	conf.API.PushURI = viper.GetString("api.push_uri")
	conf.API.StatGoURI = viper.GetString("api.stat_go_uri")
	conf.API.StatAppURI = viper.GetString("api.stat_app_uri")
	conf.API.ConfigURI = viper.GetString("api.config_uri")
	conf.API.SysStatURI = viper.GetString("api.sys_stat_uri")
	conf.API.MetricURI = viper.GetString("api.metric_uri")
	conf.API.HealthURI = viper.GetString("api.health_uri")

	// Android
	conf.Android.Enabled = viper.GetBool("android.enabled")
	conf.Android.KeyPath = viper.GetString("android.key_path")
	conf.Android.Credential = viper.GetString("android.credential")
	conf.Android.MaxRetry = viper.GetInt("android.max_retry")

	// Huawei
	conf.Huawei.Enabled = viper.GetBool("huawei.enabled")
	conf.Huawei.AppSecret = viper.GetString("huawei.appsecret")
	conf.Huawei.AppID = viper.GetString("huawei.appid")
	conf.Huawei.MaxRetry = viper.GetInt("huawei.max_retry")

	// iOS
	conf.Ios.Enabled = viper.GetBool("ios.enabled")
	conf.Ios.KeyPath = viper.GetString("ios.key_path")
	conf.Ios.KeyBase64 = viper.GetString("ios.key_base64")
	conf.Ios.KeyType = viper.GetString("ios.key_type")
	conf.Ios.Password = viper.GetString("ios.password")
	conf.Ios.Production = viper.GetBool("ios.production")
	conf.Ios.MaxConcurrentPushes = viper.GetUint("ios.max_concurrent_pushes")
	conf.Ios.MaxRetry = viper.GetInt("ios.max_retry")
	conf.Ios.KeyID = viper.GetString("ios.key_id")
	conf.Ios.TeamID = viper.GetString("ios.team_id")

	// log
	conf.Log.Format = viper.GetString("log.format")
	conf.Log.AccessLog = viper.GetString("log.access_log")
	conf.Log.AccessLevel = viper.GetString("log.access_level")
	conf.Log.ErrorLog = viper.GetString("log.error_log")
	conf.Log.ErrorLevel = viper.GetString("log.error_level")
	conf.Log.HideToken = viper.GetBool("log.hide_token")
	conf.Log.HideMessages = viper.GetBool("log.hide_messages")

	// Queue Engine
	conf.Queue.Engine = viper.GetString("queue.engine")
	conf.Queue.NSQ.Addr = viper.GetString("queue.nsq.addr")
	conf.Queue.NSQ.Topic = viper.GetString("queue.nsq.topic")
	conf.Queue.NSQ.Channel = viper.GetString("queue.nsq.channel")
	conf.Queue.NATS.Addr = viper.GetString("queue.nats.addr")
	conf.Queue.NATS.Subj = viper.GetString("queue.nats.subj")
	conf.Queue.NATS.Queue = viper.GetString("queue.nats.queue")
	conf.Queue.Redis.Addr = viper.GetString("queue.redis.addr")
	conf.Queue.Redis.StreamName = viper.GetString("queue.redis.stream_name")
	conf.Queue.Redis.Group = viper.GetString("queue.redis.group")
	conf.Queue.Redis.Consumer = viper.GetString("queue.redis.consumer")
	conf.Queue.Redis.WithTLS = viper.GetBool("queue.redis.with_tls")
	conf.Queue.Redis.Username = viper.GetString("queue.redis.username")
	conf.Queue.Redis.Password = viper.GetString("queue.redis.password")

	// Stat Engine
	conf.Stat.Engine = viper.GetString("stat.engine")
	conf.Stat.Redis.Cluster = viper.GetBool("stat.redis.cluster")
	conf.Stat.Redis.Addr = viper.GetString("stat.redis.addr")
	conf.Stat.Redis.Username = viper.GetString("stat.redis.username")
	conf.Stat.Redis.Password = viper.GetString("stat.redis.password")
	conf.Stat.Redis.DB = viper.GetInt("stat.redis.db")
	conf.Stat.BoltDB.Path = viper.GetString("stat.boltdb.path")
	conf.Stat.BoltDB.Bucket = viper.GetString("stat.boltdb.bucket")
	conf.Stat.BuntDB.Path = viper.GetString("stat.buntdb.path")
	conf.Stat.LevelDB.Path = viper.GetString("stat.leveldb.path")
	conf.Stat.BadgerDB.Path = viper.GetString("stat.badgerdb.path")

	// gRPC Server
	conf.GRPC.Enabled = viper.GetBool("grpc.enabled")
	conf.GRPC.Port = viper.GetString("grpc.port")

	if conf.Core.WorkerNum == int64(0) {
		conf.Core.WorkerNum = int64(runtime.NumCPU())
	}

	if conf.Core.QueueNum == int64(0) {
		conf.Core.QueueNum = int64(8192)
	}

	return conf, nil
}
