package internal

import "github.com/boltdb/bolt"

// Cursor that can be reversed
type Cursor struct {
	C       *bolt.Cursor
	Reverse bool
}

// First element
func (c *Cursor) First() ([]byte, []byte) {
	if c.Reverse {
		return c.C.Last()
	}

	return c.C.First()
}

// Next element
func (c *Cursor) Next() ([]byte, []byte) {
	if c.Reverse {
		return c.C.Prev()
	}

	return c.C.Next()
}

// RangeCursor that can be reversed
type RangeCursor struct {
	C         *bolt.Cursor
	Reverse   bool
	Min       []byte
	Max       []byte
	CompareFn func([]byte, []byte) int
}

// First element
func (c *RangeCursor) First() ([]byte, []byte) {
	if c.Reverse {
		return c.C.Seek(c.Max)
	}

	return c.C.Seek(c.Min)
}

// Next element
func (c *RangeCursor) Next() ([]byte, []byte) {
	if c.Reverse {
		return c.C.Prev()
	}

	return c.C.Next()
}

// Continue tells if the loop needs to continue
func (c *RangeCursor) Continue(val []byte) bool {
	if c.Reverse {
		return val != nil && c.CompareFn(val, c.Min) >= 0
	}

	return val != nil && c.CompareFn(val, c.Max) <= 0
}
