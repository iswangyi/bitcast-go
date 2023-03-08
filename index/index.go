package index

import (
	"bitcast-go/data"
	"bytes"
	"github.com/google/btree"
)

// Indexer is the interface that wraps the basic index operations.
type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}

// btreeItem is the item stored in BTree
type btreeItem struct {
	key []byte
	pos *data.LogRecordPos
}

// Less implements btree.Item interface
func (b *btreeItem) Less(than btree.Item) bool {
	return bytes.Compare(b.key, than.(*btreeItem).key) == -1
}
