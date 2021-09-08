package core

// Queue as backend
type Queue string

var (
	// LocalQueue for channel in Go
	LocalQueue Queue = "local"
	// NSQ a realtime distributed messaging platform
	NSQ Queue = "nsq"
	// NATS Connective Technology for Adaptive Edge & Distributed Systems
	NATS Queue = "nats"
)

// IsLocalQueue check is Local Queue
func IsLocalQueue(q Queue) bool {
	return q == LocalQueue
}
