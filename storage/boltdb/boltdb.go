package boltdb

import (
	"log"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/asdine/storm/v3"
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
	db     *storm.DB
	lock   sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.setBoltDB(key, s.getBoltDB(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.setBoltDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.getBoltDB(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.db, err = storm.Open(s.config.Stat.BoltDB.Path)
	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}

func (s *Storage) setBoltDB(key string, count int64) {
	err := s.db.Set(s.config.Stat.BoltDB.Bucket, key, count)
	if err != nil {
		log.Println("BoltDB set error:", err.Error())
	}
}

func (s *Storage) getBoltDB(key string) int64 {
	var count int64
	err := s.db.Get(s.config.Stat.BoltDB.Bucket, key, &count)
	if err != nil {
		log.Println("BoltDB get error:", err.Error())
	}
	return count
}
