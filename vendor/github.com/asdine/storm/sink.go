package storm

import (
	"reflect"

	"github.com/asdine/storm/index"
	"github.com/asdine/storm/q"
	"github.com/boltdb/bolt"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type item struct {
	value  *reflect.Value
	bucket *bolt.Bucket
	k      []byte
	v      []byte
}

func newSorter(node Node) *sorter {
	return &sorter{
		node:   node,
		rbTree: rbt.NewWithStringComparator(),
	}
}

type sorter struct {
	node    Node
	rbTree  *rbt.Tree
	orderBy string
	reverse bool
}

func (s *sorter) filter(snk sink, tree q.Matcher, bucket *bolt.Bucket, k, v []byte) (bool, error) {
	rsnk, ok := snk.(reflectSink)
	if !ok {
		return snk.add(&item{
			bucket: bucket,
			k:      k,
			v:      v,
		})
	}

	newElem := rsnk.elem()
	err := s.node.Codec().Unmarshal(v, newElem.Interface())
	if err != nil {
		return false, err
	}

	ok = tree == nil
	if !ok {
		ok, err = tree.Match(newElem.Interface())
		if err != nil {
			return false, err
		}
	}

	if ok {
		it := item{
			bucket: bucket,
			value:  &newElem,
			k:      k,
			v:      v,
		}

		if s.orderBy != "" {
			elm := reflect.Indirect(newElem).FieldByName(s.orderBy)
			if !elm.IsValid() {
				return false, ErrNotFound
			}
			raw, err := toBytes(elm.Interface(), s.node.Codec())
			if err != nil {
				return false, err
			}
			s.rbTree.Put(string(raw), &it)
			return false, nil
		}

		return snk.add(&it)
	}

	return false, nil
}

func (s *sorter) flush(snk sink) error {
	if s.orderBy == "" {
		return snk.flush()
	}
	s.orderBy = ""
	var err error
	var stop bool

	it := s.rbTree.Iterator()
	if s.reverse {
		it.End()
	} else {
		it.Begin()
	}
	for (s.reverse && it.Prev()) || (!s.reverse && it.Next()) {
		item := it.Value().(*item)
		stop, err = snk.add(item)
		if err != nil {
			return err
		}
		if stop {
			break
		}
	}

	return snk.flush()
}

type sink interface {
	bucketName() string
	flush() error
	add(*item) (bool, error)
}

type reflectSink interface {
	elem() reflect.Value
}

func newListSink(node Node, to interface{}) (*listSink, error) {
	ref := reflect.ValueOf(to)

	if ref.Kind() != reflect.Ptr || reflect.Indirect(ref).Kind() != reflect.Slice {
		return nil, ErrSlicePtrNeeded
	}

	sliceType := reflect.Indirect(ref).Type()
	elemType := sliceType.Elem()

	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	if elemType.Name() == "" {
		return nil, ErrNoName
	}

	return &listSink{
		node:     node,
		ref:      ref,
		isPtr:    sliceType.Elem().Kind() == reflect.Ptr,
		elemType: elemType,
		name:     elemType.Name(),
		limit:    -1,
	}, nil
}

type listSink struct {
	node     Node
	ref      reflect.Value
	results  reflect.Value
	elemType reflect.Type
	name     string
	isPtr    bool
	skip     int
	limit    int
	idx      int
}

func (l *listSink) elem() reflect.Value {
	if l.results.IsValid() && l.idx < l.results.Len() {
		return l.results.Index(l.idx).Addr()
	}
	return reflect.New(l.elemType)
}

func (l *listSink) bucketName() string {
	return l.name
}

func (l *listSink) add(i *item) (bool, error) {
	if l.limit == 0 {
		return true, nil
	}

	if l.skip > 0 {
		l.skip--
		return false, nil
	}

	if !l.results.IsValid() {
		l.results = reflect.MakeSlice(reflect.Indirect(l.ref).Type(), 0, 0)
	}

	if l.limit > 0 {
		l.limit--
	}

	if l.idx == l.results.Len() {
		if l.isPtr {
			l.results = reflect.Append(l.results, *i.value)
		} else {
			l.results = reflect.Append(l.results, reflect.Indirect(*i.value))
		}
	}

	l.idx++

	return l.limit == 0, nil
}

func (l *listSink) flush() error {
	if l.results.IsValid() && l.results.Len() > 0 {
		reflect.Indirect(l.ref).Set(l.results)
		return nil
	}

	return ErrNotFound
}

func newFirstSink(node Node, to interface{}) (*firstSink, error) {
	ref := reflect.ValueOf(to)

	if !ref.IsValid() || ref.Kind() != reflect.Ptr || ref.Elem().Kind() != reflect.Struct {
		return nil, ErrStructPtrNeeded
	}

	return &firstSink{
		node: node,
		ref:  ref,
	}, nil
}

type firstSink struct {
	node  Node
	ref   reflect.Value
	skip  int
	found bool
}

func (f *firstSink) elem() reflect.Value {
	return reflect.New(reflect.Indirect(f.ref).Type())
}

func (f *firstSink) bucketName() string {
	return reflect.Indirect(f.ref).Type().Name()
}

func (f *firstSink) add(i *item) (bool, error) {
	if f.skip > 0 {
		f.skip--
		return false, nil
	}

	reflect.Indirect(f.ref).Set(i.value.Elem())
	f.found = true
	return true, nil
}

func (f *firstSink) flush() error {
	if !f.found {
		return ErrNotFound
	}

	return nil
}

func newDeleteSink(node Node, kind interface{}) (*deleteSink, error) {
	ref := reflect.ValueOf(kind)

	if !ref.IsValid() || ref.Kind() != reflect.Ptr || ref.Elem().Kind() != reflect.Struct {
		return nil, ErrStructPtrNeeded
	}

	return &deleteSink{
		node: node,
		ref:  ref,
	}, nil
}

type deleteSink struct {
	node    Node
	ref     reflect.Value
	skip    int
	limit   int
	removed int
}

func (d *deleteSink) elem() reflect.Value {
	return reflect.New(reflect.Indirect(d.ref).Type())
}

func (d *deleteSink) bucketName() string {
	return reflect.Indirect(d.ref).Type().Name()
}

func (d *deleteSink) add(i *item) (bool, error) {
	if d.skip > 0 {
		d.skip--
		return false, nil
	}

	if d.limit > 0 {
		d.limit--
	}

	info, err := extract(&d.ref)
	if err != nil {
		return false, err
	}

	for fieldName, fieldCfg := range info.Fields {
		if fieldCfg.Index == "" {
			continue
		}
		idx, err := getIndex(i.bucket, fieldCfg.Index, fieldName)
		if err != nil {
			return false, err
		}

		err = idx.RemoveID(i.k)
		if err != nil {
			if err == index.ErrNotFound {
				return false, ErrNotFound
			}
			return false, err
		}
	}

	d.removed++
	return d.limit == 0, i.bucket.Delete(i.k)
}

func (d *deleteSink) flush() error {
	if d.removed == 0 {
		return ErrNotFound
	}

	return nil
}

func newCountSink(node Node, kind interface{}) (*countSink, error) {
	ref := reflect.ValueOf(kind)

	if !ref.IsValid() || ref.Kind() != reflect.Ptr || ref.Elem().Kind() != reflect.Struct {
		return nil, ErrStructPtrNeeded
	}

	return &countSink{
		node: node,
		ref:  ref,
	}, nil
}

type countSink struct {
	node    Node
	ref     reflect.Value
	skip    int
	limit   int
	counter int
}

func (c *countSink) elem() reflect.Value {
	return reflect.New(reflect.Indirect(c.ref).Type())
}

func (c *countSink) bucketName() string {
	return reflect.Indirect(c.ref).Type().Name()
}

func (c *countSink) add(i *item) (bool, error) {
	if c.skip > 0 {
		c.skip--
		return false, nil
	}

	if c.limit > 0 {
		c.limit--
	}

	c.counter++
	return c.limit == 0, nil
}

func (c *countSink) flush() error {
	return nil
}

func newRawSink() *rawSink {
	return &rawSink{
		limit: -1,
	}
}

type rawSink struct {
	results [][]byte
	skip    int
	limit   int
	execFn  func([]byte, []byte) error
}

func (r *rawSink) add(i *item) (bool, error) {
	if r.limit == 0 {
		return true, nil
	}

	if r.skip > 0 {
		r.skip--
		return false, nil
	}

	if r.limit > 0 {
		r.limit--
	}

	if r.execFn != nil {
		err := r.execFn(i.k, i.v)
		if err != nil {
			return false, err
		}
	} else {
		r.results = append(r.results, i.v)
	}

	return r.limit == 0, nil
}

func (r *rawSink) bucketName() string {
	return ""
}

func (r *rawSink) flush() error {
	return nil
}

func newEachSink(to interface{}) (*eachSink, error) {
	ref := reflect.ValueOf(to)

	if !ref.IsValid() || ref.Kind() != reflect.Ptr || ref.Elem().Kind() != reflect.Struct {
		return nil, ErrStructPtrNeeded
	}

	return &eachSink{
		ref: ref,
	}, nil
}

type eachSink struct {
	skip   int
	limit  int
	ref    reflect.Value
	execFn func(interface{}) error
}

func (e *eachSink) elem() reflect.Value {
	return reflect.New(reflect.Indirect(e.ref).Type())
}

func (e *eachSink) bucketName() string {
	return reflect.Indirect(e.ref).Type().Name()
}

func (e *eachSink) add(i *item) (bool, error) {
	if e.limit == 0 {
		return true, nil
	}

	if e.skip > 0 {
		e.skip--
		return false, nil
	}

	if e.limit > 0 {
		e.limit--
	}

	err := e.execFn(i.value.Interface())
	if err != nil {
		return false, err
	}

	return e.limit == 0, nil
}

func (e *eachSink) flush() error {
	return nil
}
