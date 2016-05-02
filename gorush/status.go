package gorush

import (
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v3"
	"net/http"
	"strconv"
	"sync/atomic"
)

// StatusApp is app status structure
type StatusApp struct {
	QueueMax   int           `json:"queue_max"`
	QueueUsage int           `json:"queue_usage"`
	TotalCount int64         `json:"total_count"`
	Ios        IosStatus     `json:"ios"`
	Android    AndroidStatus `json:"android"`
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

func initApp() {
	RushStatus.TotalCount = 0
	RushStatus.Ios.PushSuccess = 0
	RushStatus.Ios.PushError = 0
	RushStatus.Android.PushSuccess = 0
	RushStatus.Android.PushError = 0
}

func initRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     PushConf.Stat.Redis.Addr,
		Password: PushConf.Stat.Redis.Password,
		DB:       PushConf.Stat.Redis.DB,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		// redis server error
		LogError.Error("Can't connect redis server: " + err.Error())

		return err
	}

	RushStatus.TotalCount = getTotalCount()
	RushStatus.Ios.PushSuccess = getIosSuccess()
	RushStatus.Ios.PushError = getIosError()
	RushStatus.Android.PushSuccess = getAndroidSuccess()
	RushStatus.Android.PushError = getAndroidError()

	return nil
}

func initBoltDB() {
	RushStatus.TotalCount = getTotalCount()
	RushStatus.Ios.PushSuccess = getIosSuccess()
	RushStatus.Ios.PushError = getIosError()
	RushStatus.Android.PushSuccess = getAndroidSuccess()
	RushStatus.Android.PushError = getAndroidError()
}

// InitAppStatus for initialize app status
func InitAppStatus() {
	switch PushConf.Stat.Engine {
	case "memory":
		initApp()
	case "redis":
		initRedis()
	case "boltdb":
		initBoltDB()
	default:
		initApp()
	}

}

func getRedisInt64Result(key string, count *int64) {
	val, _ := RedisClient.Get(key).Result()
	*count, _ = strconv.ParseInt(val, 10, 64)
}

func boltdbSet(key string, count int64) {
	db, _ := storm.Open(PushConf.Stat.BoltDB.Path)
	db.Set(PushConf.Stat.BoltDB.Bucket, key, count)
	defer db.Close()
}

func boltdbGet(key string, count *int64) {
	db, _ := storm.Open(PushConf.Stat.BoltDB.Path)
	db.Get(PushConf.Stat.BoltDB.Bucket, key, count)
	defer db.Close()
}

func addTotalCount(count int64) {
	switch PushConf.Stat.Engine {
	case "memory":
		atomic.AddInt64(&RushStatus.TotalCount, count)
	case "redis":
		RedisClient.Set(TotalCountKey, strconv.Itoa(int(count)), 0)
	case "boltdb":
		boltdbSet(TotalCountKey, count)
	default:
		atomic.AddInt64(&RushStatus.TotalCount, count)
	}
}

func addIosSuccess(count int64) {
	switch PushConf.Stat.Engine {
	case "memory":
		atomic.AddInt64(&RushStatus.Ios.PushSuccess, count)
	case "redis":
		RedisClient.Set(IosSuccessKey, strconv.Itoa(int(count)), 0)
	case "boltdb":
		boltdbSet(IosSuccessKey, count)
	default:
		atomic.AddInt64(&RushStatus.Ios.PushSuccess, count)
	}
}

func addIosError(count int64) {
	switch PushConf.Stat.Engine {
	case "memory":
		atomic.AddInt64(&RushStatus.Ios.PushError, count)
	case "redis":
		RedisClient.Set(IosErrorKey, strconv.Itoa(int(count)), 0)
	case "boltdb":
		boltdbSet(IosErrorKey, count)
	default:
		atomic.AddInt64(&RushStatus.Ios.PushError, count)
	}
}

func addAndroidSuccess(count int64) {
	switch PushConf.Stat.Engine {
	case "memory":
		atomic.AddInt64(&RushStatus.Android.PushSuccess, count)
	case "redis":
		RedisClient.Set(AndroidSuccessKey, strconv.Itoa(int(count)), 0)
	case "boltdb":
		boltdbSet(AndroidSuccessKey, count)
	default:
		atomic.AddInt64(&RushStatus.Android.PushSuccess, count)
	}
}

func addAndroidError(count int64) {
	switch PushConf.Stat.Engine {
	case "memory":
		atomic.AddInt64(&RushStatus.Android.PushError, count)
	case "redis":
		RedisClient.Set(AndroidErrorKey, strconv.Itoa(int(count)), 0)
	case "boltdb":
		boltdbSet(AndroidErrorKey, count)
	default:
		atomic.AddInt64(&RushStatus.Android.PushError, count)
	}
}

func getTotalCount() int64 {
	var count int64
	switch PushConf.Stat.Engine {
	case "memory":
		count = atomic.LoadInt64(&RushStatus.TotalCount)
	case "redis":
		getRedisInt64Result(TotalCountKey, &count)
	case "boltdb":
		boltdbGet(TotalCountKey, &count)
	default:
		count = atomic.LoadInt64(&RushStatus.TotalCount)
	}

	return count
}

func getIosSuccess() int64 {
	var count int64
	switch PushConf.Stat.Engine {
	case "memory":
		count = atomic.LoadInt64(&RushStatus.Ios.PushSuccess)
	case "redis":
		getRedisInt64Result(IosSuccessKey, &count)
	case "boltdb":
		boltdbGet(IosSuccessKey, &count)
	default:
		count = atomic.LoadInt64(&RushStatus.Ios.PushSuccess)
	}

	return count
}

func getIosError() int64 {
	var count int64
	switch PushConf.Stat.Engine {
	case "memory":
		count = atomic.LoadInt64(&RushStatus.Ios.PushError)
	case "redis":
		getRedisInt64Result(IosErrorKey, &count)
	case "boltdb":
		boltdbGet(IosErrorKey, &count)
	default:
		count = atomic.LoadInt64(&RushStatus.Ios.PushError)
	}

	return count
}

func getAndroidSuccess() int64 {
	var count int64
	switch PushConf.Stat.Engine {
	case "memory":
		count = atomic.LoadInt64(&RushStatus.Android.PushSuccess)
	case "redis":
		getRedisInt64Result(AndroidSuccessKey, &count)
	case "boltdb":
		boltdbGet(AndroidSuccessKey, &count)
	default:
		count = atomic.LoadInt64(&RushStatus.Android.PushSuccess)
	}

	return count
}

func getAndroidError() int64 {
	var count int64
	switch PushConf.Stat.Engine {
	case "memory":
		count = atomic.LoadInt64(&RushStatus.Android.PushError)
	case "redis":
		getRedisInt64Result(AndroidErrorKey, &count)
	case "boltdb":
		boltdbGet(AndroidErrorKey, &count)
	default:
		count = atomic.LoadInt64(&RushStatus.Android.PushError)
	}

	return count
}

func appStatusHandler(c *gin.Context) {
	result := StatusApp{}

	result.QueueMax = cap(QueueNotification)
	result.QueueUsage = len(QueueNotification)
	result.TotalCount = getTotalCount()
	result.Ios.PushSuccess = getIosSuccess()
	result.Ios.PushError = getIosError()
	result.Android.PushSuccess = getAndroidSuccess()
	result.Android.PushError = getAndroidError()

	c.JSON(http.StatusOK, result)
}
