package status

import (
	"errors"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/storage"
	"github.com/appleboy/gorush/storage/badger"
	"github.com/appleboy/gorush/storage/boltdb"
	"github.com/appleboy/gorush/storage/buntdb"
	"github.com/appleboy/gorush/storage/leveldb"
	"github.com/appleboy/gorush/storage/memory"
	"github.com/appleboy/gorush/storage/redis"

	"github.com/thoas/stats"
)

// Stats provide response time, status code count, etc.
var Stats *stats.Stats

// StatStorage implements the storage interface
var StatStorage storage.Storage

// App is status structure
type App struct {
	Version    string        `json:"version"`
	QueueMax   int           `json:"queue_max"`
	QueueUsage int           `json:"queue_usage"`
	TotalCount int64         `json:"total_count"`
	Ios        IosStatus     `json:"ios"`
	Android    AndroidStatus `json:"android"`
	Huawei     HuaweiStatus  `json:"huawei"`
}

// AndroidStatus is android structure
type AndroidStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// IosStatus is iOS structure
type IosStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// HuaweiStatus is huawei structure
type HuaweiStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// InitAppStatus for initialize app status
func InitAppStatus(conf *config.ConfYaml) error {
	logx.LogAccess.Info("Init App Status Engine as ", conf.Stat.Engine)
	switch conf.Stat.Engine {
	case "memory":
		StatStorage = memory.New()
	case "redis":
		StatStorage = redis.New(conf)
	case "boltdb":
		StatStorage = boltdb.New(conf)
	case "buntdb":
		StatStorage = buntdb.New(conf)
	case "leveldb":
		StatStorage = leveldb.New(conf)
	case "badger":
		StatStorage = badger.New(conf)
	default:
		logx.LogError.Error("storage error: can't find storage driver")
		return errors.New("can't find storage driver")
	}

	if err := StatStorage.Init(); err != nil {
		logx.LogError.Error("storage error: " + err.Error())

		return err
	}

	Stats = stats.New()

	return nil
}
