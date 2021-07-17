package nsq

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/queue"

	"github.com/nsqio/go-nsq"
)

var _ queue.Worker = (*Worker)(nil)

// Option for queue system
type Option func(*Worker)

// Worker for NSQ
type Worker struct {
	q       *nsq.Consumer
	p       *nsq.Producer
	once    sync.Once
	addr    string
	topic   string
	channel string
}

// WithAddr setup the addr of NSQ
func WithAddr(addr string) Option {
	return func(w *Worker) {
		w.addr = addr
	}
}

// WithTopic setup the topic of NSQ
func WithTopic(topic string) Option {
	return func(w *Worker) {
		w.topic = topic
	}
}

// WithChannel setup the channel of NSQ
func WithChannel(channel string) Option {
	return func(w *Worker) {
		w.channel = channel
	}
}

// NewWorker for struc
func NewWorker(opts ...Option) *Worker {
	w := &Worker{
		addr:    "127.0.0.1:4150",
		topic:   "gorush",
		channel: "ch",
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(w)
	}

	cfg := nsq.NewConfig()
	q, err := nsq.NewConsumer(w.topic, w.channel, cfg)
	if err != nil {
		panic(err)
	}
	w.q = q

	p, err := nsq.NewProducer(w.addr, cfg)
	if err != nil {
		panic(err)
	}
	w.p = p

	return w
}

// BeforeRun run script before start worker
func (s *Worker) BeforeRun() error {
	return nil
}

// AfterRun run script after start worker
func (s *Worker) AfterRun() error {
	s.once.Do(func() {
		time.Sleep(100 * time.Millisecond)
		err := s.q.ConnectToNSQD(s.addr)
		if err != nil {
			panic("Could not connect nsq server: " + err.Error())
		}
	})

	return nil
}

// Run start the worker
func (s *Worker) Run(quit chan struct{}) error {
	wg := &sync.WaitGroup{}
	s.q.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		wg.Add(1)
		defer wg.Done()
		var notification gorush.PushNotification
		if err := json.Unmarshal(msg.Body, &notification); err != nil {
			return err
		}
		gorush.SendNotification(notification)
		time.Sleep(10 * time.Second)
		return nil
	}))

	select {
	case <-quit:
	}

	wg.Wait()
	return nil
}

// Shutdown worker
func (s *Worker) Shutdown() error {
	s.q.Stop()
	s.p.Stop()
	return nil
}

// Capacity for channel
func (s *Worker) Capacity() int {
	return 0
}

// Usage for count of channel usage
func (s *Worker) Usage() int {
	return 0
}

// Queue send notification to queue
func (s *Worker) Queue(job interface{}) error {
	v, ok := job.(gorush.PushNotification)
	if !ok {
		return errors.New("wrong type of job")
	}
	err := s.p.Publish(s.topic, v.Bytes())
	if err != nil {
		return err
	}

	return nil
}
