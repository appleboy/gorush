package app

import (
	"fmt"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/notify"

	"github.com/golang-queue/nats"
	"github.com/golang-queue/nsq"
	"github.com/golang-queue/queue"
	qcore "github.com/golang-queue/queue/core"
	redisdb "github.com/golang-queue/redisdb-stream"
)

// NewQueueWorker creates a queue worker based on the configured queue engine.
// Supported engines: local, nsq, nats, redis.
func NewQueueWorker(cfg *config.ConfYaml) (qcore.Worker, error) {
	switch core.Queue(cfg.Queue.Engine) {
	case core.LocalQueue:
		return queue.NewRing(
			queue.WithQueueSize(int(cfg.Core.QueueNum)),
			queue.WithFn(notify.Run(cfg)),
			queue.WithLogger(logx.QueueLogger()),
		), nil

	case core.NSQ:
		return nsq.NewWorker(
			nsq.WithAddr(cfg.Queue.NSQ.Addr),
			nsq.WithTopic(cfg.Queue.NSQ.Topic),
			nsq.WithChannel(cfg.Queue.NSQ.Channel),
			nsq.WithMaxInFlight(int(cfg.Core.WorkerNum)),
			nsq.WithRunFunc(notify.Run(cfg)),
			nsq.WithLogger(logx.QueueLogger()),
		), nil

	case core.NATS:
		return nats.NewWorker(
			nats.WithAddr(cfg.Queue.NATS.Addr),
			nats.WithSubj(cfg.Queue.NATS.Subj),
			nats.WithQueue(cfg.Queue.NATS.Queue),
			nats.WithRunFunc(notify.Run(cfg)),
			nats.WithLogger(logx.QueueLogger()),
		), nil

	case core.Redis:
		opts := []redisdb.Option{
			redisdb.WithAddr(cfg.Queue.Redis.Addr),
			redisdb.WithUsername(cfg.Queue.Redis.Username),
			redisdb.WithPassword(cfg.Queue.Redis.Password),
			redisdb.WithDB(cfg.Queue.Redis.DB),
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
		return redisdb.NewWorker(opts...), nil

	default:
		return nil, fmt.Errorf("unsupported queue engine: %s", cfg.Queue.Engine)
	}
}

// NewQueuePool creates a queue pool with the configured number of workers.
func NewQueuePool(cfg *config.ConfYaml, w qcore.Worker) *queue.Queue {
	return queue.NewPool(
		cfg.Core.WorkerNum,
		queue.WithWorker(w),
		queue.WithLogger(logx.QueueLogger()),
	)
}
