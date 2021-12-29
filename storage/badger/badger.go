package badger

import (
	"log"
	"os"
	"strconv"

	"github.com/appleboy/gorush/config"
	"github.com/dgraph-io/badger/v3"
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
	opts   badger.Options
	name   string
	db     *badger.DB
}

func (s *Storage) Add(key string, count int64) {
	s.Set(key, s.Get(key)+count)
}

func (s *Storage) Set(key string, count int64) {
	s.setBadger(key, count)
}

func (s *Storage) Get(key string) int64 {
	var count int64
	s.getBadger(key, &count)

	return count
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.name = "badger"
	dbPath := s.config.Stat.BadgerDB.Path
	if dbPath == "" {
		dbPath = os.TempDir() + "badger"
	}
	s.opts = badger.DefaultOptions(dbPath)

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

func (s *Storage) getBadger(key string, count *int64) {
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		dst := []byte{}
		val, err := item.ValueCopy(dst)
		if err != nil {
			return err
		}

		i, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return err
		}

		*count = i

		return nil
	})
	if err != nil {
		log.Println(s.name, "get error:", err.Error())
	}
}
