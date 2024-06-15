package badger

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/appleboy/gorush/core"

	"github.com/dgraph-io/badger/v4"
)

var _ core.Storage = (*Storage)(nil)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(dbPath string) *Storage {
	return &Storage{
		dbPath: dbPath,
	}
}

// Storage is interface structure
type Storage struct {
	dbPath string
	opts   badger.Options
	name   string
	db     *badger.DB

	sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setBadger(key, s.getBadger(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setBadger(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.RLock()
	defer s.RUnlock()
	return s.getBadger(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.name = "badger"
	if s.dbPath == "" {
		s.dbPath = os.TempDir() + "badger"
	}
	s.opts = badger.DefaultOptions(s.dbPath)

	s.db, err = badger.Open(s.opts)

	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}

func (s *Storage) setBadger(key string, count int64) {
	err := s.db.Update(func(txn *badger.Txn) error {
		value := strconv.FormatInt(count, 10)
		return txn.Set([]byte(key), []byte(value))
	})
	if err != nil {
		log.Println(s.name, "update error:", err.Error())
	}
}

func (s *Storage) getBadger(key string) int64 {
	var count int64
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		var dst []byte
		val, err := item.ValueCopy(dst)
		if err != nil {
			return err
		}

		count, err = strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Println(s.name, "get error:", err.Error())
	}
	return count
}
