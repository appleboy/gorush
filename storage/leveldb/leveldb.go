package leveldb

import (
	"fmt"
	"strconv"

	"github.com/appleboy/gorush/config"
	"github.com/syndtr/goleveldb/leveldb"
)

func (s *Storage) setLevelDB(key string, count int64) {
	value := fmt.Sprintf("%d", count)
	_ = s.db.Put([]byte(key), []byte(value), nil)
}

func (s *Storage) getLevelDB(key string, count *int64) {
	data, _ := s.db.Get([]byte(key), nil)
	*count, _ = strconv.ParseInt(string(data), 10, 64)
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
}

func (s *Storage) Add(key string, count int64) {
	s.Set(key, s.Get(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.setLevelDB(key, count)
}

func (s *Storage) Get(key string) int64 {
	var count int64
	s.getLevelDB(key, &count)

	return count
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
