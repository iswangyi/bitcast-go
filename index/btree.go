package index

import (
	"bitcast-go/data"
	"github.com/google/btree"
	"sync"
)

type BTree struct {
	tree *btree.BTree
	//Btree is not thread safe for write operation, so we need a lock
	//Btree is thread safe for read operation
	lock *sync.RWMutex
}

// Put Btree implements Indexer interface
// Put 向索引中插入对应的文件位置信息
func (B *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	B.lock.Lock()
	defer B.lock.Unlock()
	B.tree.ReplaceOrInsert(&btreeItem{key, pos})

	return true
}

// Get Btree implements Indexer interface
// Get 从索引中获取对应的文件位置信息
func (B *BTree) Get(key []byte) *data.LogRecordPos {
	B.lock.RLock()
	defer B.lock.RUnlock()
	item := B.tree.Get(&btreeItem{key, nil})
	if item == nil {
		return nil
	}
	return item.(*btreeItem).pos
}

// Delete Btree implements Indexer interface
// Delete 从索引中删除对应的文件位置信息
func (B *BTree) Delete(key []byte) bool {
	B.lock.Lock()
	defer B.lock.Unlock()
	OldItem := B.tree.Delete(&btreeItem{key, nil})
	if OldItem == nil {
		return false
	}
	return true
}

// NewBTree returns a new BTree
func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: &sync.RWMutex{},
	}
}
