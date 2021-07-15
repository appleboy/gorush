package queue

import "github.com/appleboy/gorush/config"

// Worker interface
type Worker interface {
	Run(chan struct{})
	Stop()
	Enqueue(job interface{}) error
	Capacity() int
	Usage() int
	Config(config.ConfYaml)
}
