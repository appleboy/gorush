package boltdb

import (
	"log"
	"os"
	"sync"

	"github.com/appleboy/gorush/core"

	"github.com/asdine/storm/v3"
)

var _ core.Storage = (*Storage)(nil)

// New func implements the storage interface for gorush (https://github.com/appleboy/gorush)
func New(dbPath, bucket string) *Storage {
	return &Storage{
		dbPath: dbPath,
		bucket: bucket,
	}
}

// Storage is interface structure
type Storage struct {
	dbPath string
	bucket string
	db     *storm.DB
	sync.RWMutex
}

func (s *Storage) Add(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setBoltDB(key, s.getBoltDB(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.Lock()
	defer s.Unlock()
	s.setBoltDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	s.RLock()
	defer s.RUnlock()
	return s.getBoltDB(key)
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	if s.dbPath == "" {
		s.dbPath = os.TempDir() + "boltdb.db"
	}
	s.db, err = storm.Open(s.dbPath)
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
	err := s.db.Set(s.bucket, key, count)
	if err != nil {
		log.Println("BoltDB set error:", err.Error())
	}
}

func (s *Storage) getBoltDB(key string) int64 {
	var count int64
	err := s.db.Get(s.bucket, key, &count)
	if err != nil {
		log.Println("BoltDB get error:", err.Error())
	}
	return count
}
