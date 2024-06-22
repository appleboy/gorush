package memory

import (
	"sync"

	"github.com/appleboy/gorush/core"

	"go.uber.org/atomic"
)

var _ core.Storage = (*Storage)(nil)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New() *Storage {
	return &Storage{}
}

// Storage is interface structure
type Storage struct {
	mem sync.Map
}

func (s *Storage) getValueBtKey(key string) *atomic.Int64 {
	if val, ok := s.mem.Load(key); ok {
		return val.(*atomic.Int64)
	}
	val := atomic.NewInt64(0)
	s.mem.Store(key, val)
	return val
}

func (s *Storage) Add(key string, count int64) {
	s.getValueBtKey(key).Add(count)
}

func (s *Storage) Set(key string, count int64) {
	s.getValueBtKey(key).Store(count)
}

func (s *Storage) Get(key string) int64 {
	return s.getValueBtKey(key).Load()
}

// Init client storage.
func (*Storage) Init() error {
	return nil
}

// Close the storage connection
func (*Storage) Close() error {
	return nil
}
