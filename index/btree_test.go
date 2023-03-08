package index

import (
	"bitcast-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTree_Put(t *testing.T) {
	bt := NewBTree()
	res1 := bt.Put([]byte("a"), &data.LogRecordPos{Offset: 1, Fid: 1})
	assert.True(t, res1)

	res2 := bt.Put(nil, &data.LogRecordPos{Offset: 2, Fid: 2})
	assert.True(t, res2)

	res3 := bt.Put(nil, nil)
	assert.True(t, res3)
}

func TestBTree_Delete(t *testing.T) {
	bt := NewBTree()
	bt.Put([]byte("a"), &data.LogRecordPos{Offset: 1, Fid: 1})
	bt.Delete([]byte("a"))
}

func TestBtree_Get(t *testing.T) {
	bt := NewBTree()
	bt.Put([]byte("a"), &data.LogRecordPos{Offset: 1, Fid: 1})
	bt.Get([]byte("a"))
	bt.Put(nil, &data.LogRecordPos{Offset: 1, Fid: 1})
}
