package index

import (
	"bitcask/data"
	"bytes"

	"github.com/google/btree"
)

// Abstract index interface, subsequently if want to access other data structures, the direct implementation of this interface can be.
type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) *data.LogRecordPos // Put 向索引中存储 key 对应的数据位置信息
	Get(key []byte) *data.LogRecordPos                         // Get 根据 key 取出对应的索引位置信息
	Delete(key []byte) (*data.LogRecordPos, bool)              // Delete 根据 key 删除对应的索引位置信息
	Size() int                                                 // Size 索引中的数据量
	Iterator(reverse bool) Iterator                            // Iterator 索引迭代器
	Close() error                                              // Close 关闭索引
}

type IndexType = int8

const (
	Btree  IndexType = iota + 1 // Btree 索引
	ART                         // 自适应基数树索引
	BPTree                      // B+ 树索引
)

// NewIndexer 根据类型初始化索引
func NewIndexer(typ IndexType, dirPath string, sync bool) Indexer {
	switch typ {
	case Btree:
		return NewBtree()
	case ART:
		return NewART()
	case BPTree:
		return NewBPlusTree(dirPath, sync)
	default:
		panic("unsuppoted index type")
	}
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

// Iterator 通用索引迭代器
type Iterator interface {
	Rewind()                   // Rewind 重新回到迭代器的起点，即第一个数据
	Seek(key []byte)           // Seek 根据传入的 key 查找到第一个大于（或小于）等于的目标 key，根据从这个 key开始遍历
	Next()                     // Next 跳转到下一个 key
	Valid() bool               // Valid 是否有效，即是否已经遍历完了所有的 key，用于退出遍历
	Key() []byte               // Key 当前遍历位置的 key 的数据
	Value() *data.LogRecordPos // Value 当前遍历位置的 value 数据
	Close()                    // Close关闭迭代器，释放相应资源
}
