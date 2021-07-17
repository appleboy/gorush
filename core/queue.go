package core

type Queue string

var (
	LocalQueue Queue = "local"
	NSQ        Queue = "nsq"
	NATS       Queue = "nats"
)

// IsLocalQueue check is Local Queue
func IsLocalQueue(q Queue) bool {
	return q == LocalQueue
}
