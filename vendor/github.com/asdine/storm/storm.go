package storm

import (
	"bytes"
	"encoding/binary"
	"os"
	"time"

	"github.com/asdine/storm/codec"
	"github.com/asdine/storm/codec/json"
	"github.com/asdine/storm/index"
	"github.com/asdine/storm/q"
	"github.com/boltdb/bolt"
)

const (
	dbinfo         = "__storm_db"
	metadataBucket = "__storm_metadata"
)

// Defaults to json
var defaultCodec = json.Codec

// Open opens a database at the given path with optional Storm options.
func Open(path string, stormOptions ...func(*DB) error) (*DB, error) {
	var err error

	s := &DB{
		Path:  path,
		codec: defaultCodec,
	}

	for _, option := range stormOptions {
		if err = option(s); err != nil {
			return nil, err
		}
	}

	if s.boltMode == 0 {
		s.boltMode = 0600
	}

	if s.boltOptions == nil {
		s.boltOptions = &bolt.Options{Timeout: 1 * time.Second}
	}

	s.root = &node{s: s, rootBucket: s.rootBucket, codec: s.codec, batchMode: s.batchMode}

	// skip if UseDB option is used
	if s.Bolt == nil {
		s.Bolt, err = bolt.Open(path, s.boltMode, s.boltOptions)
		if err != nil {
			return nil, err
		}

		err = s.checkVersion()
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

// DB is the wrapper around BoltDB. It contains an instance of BoltDB and uses it to perform all the
// needed operations
type DB struct {
	// Path of the database file
	Path string

	// Handles encoding and decoding of objects
	codec codec.MarshalUnmarshaler

	// Bolt is still easily accessible
	Bolt *bolt.DB

	// Bolt file mode
	boltMode os.FileMode

	// Bolt options
	boltOptions *bolt.Options

	// Enable auto increment on empty integer fields
	autoIncrement bool

	// The root node that points to the root bucket.
	root *node

	// The root bucket name
	rootBucket []string

	// Enable batch mode for read-write transaction, instead of update mode
	batchMode bool
}

// From returns a new Storm node with a new bucket root.
// All DB operations on the new node will be executed relative to the given
// bucket.
func (s *DB) From(root ...string) Node {
	newNode := *s.root
	newNode.rootBucket = root
	return &newNode
}

// WithTransaction returns a New Storm node that will use the given transaction.
func (s *DB) WithTransaction(tx *bolt.Tx) Node {
	return s.root.WithTransaction(tx)
}

// Bucket returns the root bucket name as a slice.
// In the normal, simple case this will be empty.
func (s *DB) Bucket() []string {
	return s.root.Bucket()
}

// Close the database
func (s *DB) Close() error {
	return s.Bolt.Close()
}

// Codec returns the EncodeDecoder used by this instance of Storm
func (s *DB) Codec() codec.MarshalUnmarshaler {
	return s.codec
}

// WithCodec returns a New Storm Node that will use the given Codec.
func (s *DB) WithCodec(codec codec.MarshalUnmarshaler) Node {
	n := s.From().(*node)
	n.codec = codec
	return n
}

// WithBatch returns a new Storm Node with the batch mode enabled.
func (s *DB) WithBatch(enabled bool) Node {
	n := s.From().(*node)
	n.batchMode = enabled
	return n
}

// Get a value from a bucket
func (s *DB) Get(bucketName string, key interface{}, to interface{}) error {
	return s.root.Get(bucketName, key, to)
}

// Set a key/value pair into a bucket
func (s *DB) Set(bucketName string, key interface{}, value interface{}) error {
	return s.root.Set(bucketName, key, value)
}

// Delete deletes a key from a bucket
func (s *DB) Delete(bucketName string, key interface{}) error {
	return s.root.Delete(bucketName, key)
}

// GetBytes gets a raw value from a bucket.
func (s *DB) GetBytes(bucketName string, key interface{}) ([]byte, error) {
	return s.root.GetBytes(bucketName, key)
}

// SetBytes sets a raw value into a bucket.
func (s *DB) SetBytes(bucketName string, key interface{}, value []byte) error {
	return s.root.SetBytes(bucketName, key, value)
}

// Save a structure
func (s *DB) Save(data interface{}) error {
	return s.root.Save(data)
}

// PrefixScan scans the root buckets for keys matching the given prefix.
func (s *DB) PrefixScan(prefix string) []Node {
	return s.root.PrefixScan(prefix)
}

// RangeScan scans the root buckets over a range such as a sortable time range.
func (s *DB) RangeScan(min, max string) []Node {
	return s.root.RangeScan(min, max)
}

// Select a list of records that match a list of matchers. Doesn't use indexes.
func (s *DB) Select(matchers ...q.Matcher) Query {
	return s.root.Select(matchers...)
}

// Range returns one or more records by the specified index within the specified range
func (s *DB) Range(fieldName string, min, max, to interface{}, options ...func(*index.Options)) error {
	return s.root.Range(fieldName, min, max, to, options...)
}

// AllByIndex gets all the records of a bucket that are indexed in the specified index
func (s *DB) AllByIndex(fieldName string, to interface{}, options ...func(*index.Options)) error {
	return s.root.AllByIndex(fieldName, to, options...)
}

// All get all the records of a bucket
func (s *DB) All(to interface{}, options ...func(*index.Options)) error {
	return s.root.All(to, options...)
}

// Count counts all the records of a bucket
func (s *DB) Count(data interface{}) (int, error) {
	return s.root.Count(data)
}

// DeleteStruct deletes a structure from the associated bucket
func (s *DB) DeleteStruct(data interface{}) error {
	return s.root.DeleteStruct(data)
}

// Remove deletes a structure from the associated bucket
// Deprecated: Use DeleteStruct instead.
func (s *DB) Remove(data interface{}) error {
	return s.root.DeleteStruct(data)
}

// Drop a bucket
func (s *DB) Drop(data interface{}) error {
	return s.root.Drop(data)
}

// Find returns one or more records by the specified index
func (s *DB) Find(fieldName string, value interface{}, to interface{}, options ...func(q *index.Options)) error {
	return s.root.Find(fieldName, value, to, options...)
}

// Init creates the indexes and buckets for a given structure
func (s *DB) Init(data interface{}) error {
	return s.root.Init(data)
}

// ReIndex rebuilds all the indexes of a bucket
func (s *DB) ReIndex(data interface{}) error {
	return s.root.ReIndex(data)
}

// One returns one record by the specified index
func (s *DB) One(fieldName string, value interface{}, to interface{}) error {
	return s.root.One(fieldName, value, to)
}

// Begin starts a new transaction.
func (s *DB) Begin(writable bool) (Node, error) {
	return s.root.Begin(writable)
}

// Rollback closes the transaction and ignores all previous updates.
func (s *DB) Rollback() error {
	return s.root.Rollback()
}

// Commit writes all changes to disk.
func (s *DB) Commit() error {
	return s.root.Rollback()
}

// Update a structure
func (s *DB) Update(data interface{}) error {
	return s.root.Update(data)
}

// UpdateField updates a single field
func (s *DB) UpdateField(data interface{}, fieldName string, value interface{}) error {
	return s.root.UpdateField(data, fieldName, value)
}

// CreateBucketIfNotExists creates the bucket below the current node if it doesn't
// already exist.
func (s *DB) CreateBucketIfNotExists(tx *bolt.Tx, bucket string) (*bolt.Bucket, error) {
	return s.root.CreateBucketIfNotExists(tx, bucket)
}

// GetBucket returns the given bucket below the current node.
func (s *DB) GetBucket(tx *bolt.Tx, children ...string) *bolt.Bucket {
	return s.root.GetBucket(tx, children...)
}

func (s *DB) checkVersion() error {
	var v string
	err := s.Get(dbinfo, "version", &v)
	if err != nil && err != ErrNotFound {
		return err
	}

	// for now, we only set the current version if it doesn't exist or if v0.5.0
	if v == "" || v == "0.5.0" || v == "0.6.0" {
		return s.Set(dbinfo, "version", Version)
	}

	return nil
}

// toBytes turns an interface into a slice of bytes
func toBytes(key interface{}, codec codec.MarshalUnmarshaler) ([]byte, error) {
	if key == nil {
		return nil, nil
	}
	switch t := key.(type) {
	case []byte:
		return t, nil
	case string:
		return []byte(t), nil
	case int:
		return numbertob(int64(t))
	case uint:
		return numbertob(uint64(t))
	case int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return numbertob(t)
	default:
		return codec.Marshal(key)
	}
}

func numbertob(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func numberfromb(raw []byte) (int64, error) {
	r := bytes.NewReader(raw)
	var to int64
	err := binary.Read(r, binary.BigEndian, &to)
	if err != nil {
		return 0, err
	}
	return to, nil
}
