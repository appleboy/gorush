package buntdb

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/tidwall/buntdb"
)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config *config.ConfYaml
	db     *buntdb.DB
	lock   sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.setBuntDB(key, s.getBuntDB(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.setBuntDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.getBuntDB(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.db, err = buntdb.Open(s.config.Stat.BuntDB.Path)
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
