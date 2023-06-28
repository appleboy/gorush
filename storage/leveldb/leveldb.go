package leveldb

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/syndtr/goleveldb/leveldb"
)

func (s *Storage) setLevelDB(key string, count int64) {
	value := fmt.Sprintf("%d", count)
	_ = s.db.Put([]byte(key), []byte(value), nil)
}

func (s *Storage) getLevelDB(key string) int64 {
	data, _ := s.db.Get([]byte(key), nil)
	count, _ := strconv.ParseInt(string(data), 10, 64)
	return count
}

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config *config.ConfYaml
	db     *leveldb.DB
	lock   sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.setLevelDB(key, s.getLevelDB(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.setLevelDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.getLevelDB(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.db, err = leveldb.OpenFile(s.config.Stat.LevelDB.Path, nil)
	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}
