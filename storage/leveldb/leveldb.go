package leveldb

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/appleboy/gorush/core"

	"github.com/syndtr/goleveldb/leveldb"
)

var _ core.Storage = (*Storage)(nil)

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
func New(dbPath string) *Storage {
	return &Storage{
		dbPath: dbPath,
	}
}

// Storage is interface structure
type Storage struct {
	dbPath string
	db     *leveldb.DB
	sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setLevelDB(key, s.getLevelDB(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setLevelDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.RLock()
	defer s.RUnlock()
	return s.getLevelDB(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	if s.dbPath == "" {
		s.dbPath = os.TempDir() + "leveldb.db"
	}
	s.db, err = leveldb.OpenFile(s.dbPath, nil)
	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}
