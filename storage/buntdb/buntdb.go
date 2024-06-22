package buntdb

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/appleboy/gorush/core"

	"github.com/tidwall/buntdb"
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
	db     *buntdb.DB
	sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setBuntDB(key, s.getBuntDB(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setBuntDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.RLock()
	defer s.RUnlock()
	return s.getBuntDB(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	if s.dbPath == "" {
		s.dbPath = os.TempDir() + "buntdb.db"
	}
	s.db, err = buntdb.Open(s.dbPath)
	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}

func (s *Storage) setBuntDB(key string, count int64) {
	err := s.db.Update(func(tx *buntdb.Tx) error {
		if _, _, err := tx.Set(key, fmt.Sprintf("%d", count), nil); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("BuntDB update error:", err.Error())
	}
}

func (s *Storage) getBuntDB(key string) int64 {
	var count int64
	err := s.db.View(func(tx *buntdb.Tx) error {
		val, _ := tx.Get(key)
		count, _ = strconv.ParseInt(val, 10, 64)
		return nil
	})
	if err != nil {
		log.Println("BuntDB get error:", err.Error())
	}

	return count
}
