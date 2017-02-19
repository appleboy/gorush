package storm

import (
	"os"

	"github.com/asdine/storm/codec"
	"github.com/asdine/storm/index"
	"github.com/boltdb/bolt"
)

// BoltOptions used to pass options to BoltDB.
func BoltOptions(mode os.FileMode, options *bolt.Options) func(*DB) error {
	return func(d *DB) error {
		d.boltMode = mode
		d.boltOptions = options
		return nil
	}
}

// Codec used to set a custom encoder and decoder. The default is JSON.
func Codec(c codec.MarshalUnmarshaler) func(*DB) error {
	return func(d *DB) error {
		d.codec = c
		return nil
	}
}

// Batch enables the use of batch instead of update for read-write transactions.
func Batch() func(*DB) error {
	return func(d *DB) error {
		d.batchMode = true
		return nil
	}
}

// AutoIncrement used to enable bolt.NextSequence on empty integer ids.
// Deprecated: Set the increment tag to the id field instead.
func AutoIncrement() func(*DB) error {
	return func(d *DB) error {
		d.autoIncrement = true
		return nil
	}
}

// Root used to set the root bucket. See also the From method.
func Root(root ...string) func(*DB) error {
	return func(d *DB) error {
		d.rootBucket = root
		return nil
	}
}

// UseDB allow Storm to use an existing open Bolt.DB.
// Warning: storm.DB.Close() will close the bolt.DB instance.
func UseDB(b *bolt.DB) func(*DB) error {
	return func(d *DB) error {
		d.Path = b.Path()
		d.Bolt = b
		return nil
	}
}

// Limit sets the maximum number of records to return
func Limit(limit int) func(*index.Options) {
	return func(opts *index.Options) {
		opts.Limit = limit
	}
}

// Skip sets the number of records to skip
func Skip(offset int) func(*index.Options) {
	return func(opts *index.Options) {
		opts.Skip = offset
	}
}

// Reverse will return the results in descending order
func Reverse() func(*index.Options) {
	return func(opts *index.Options) {
		opts.Reverse = true
	}
}
