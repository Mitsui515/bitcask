package bitcask

import "os"

// 配置项
type Options struct {
	DirPath            string      // 数据库数据目录
	DataFileSize       int64       // 数据文件的大小
	SyncWrites         bool        // 每次写入数据是否持久化
	BytesPerSync       uint        // 累计写道多少字节后进行持久化
	IndexType          IndexerType // 索引类型
	MMapAtStartup      bool        // 启动时是否使用 MMap 加载
	DataFileMergeRatio float32     // 数据文件合并的阈值
}

// 索引迭代器配置项
type IteratorOptions struct {
	Prefix  []byte // 遍历前缀为指定值的 Key，默认为空
	Reverse bool   // 是否反向遍历，默认 false 是正向
}

// 皮凉鞋配置项
type WriteBatchOptions struct {
	MaxBatchNum uint // 一个批次当中最大的数据量
	SyncWrites  bool // 提交事务时是否进行持久化
}

type IndexerType = int8

const (
	BTree     IndexerType = iota + 1 // BTree 索引
	ART                              // ART 自适应基数树索引
	BPlusTree                        // BPlusTree B+Tree 索引，将索引存储到磁盘上
)

var DefaultOptions = Options{
	DirPath:            os.TempDir(),
	DataFileSize:       256 * 1024 * 1024, // 256MB
	SyncWrites:         false,
	BytesPerSync:       0,
	IndexType:          BTree,
	MMapAtStartup:      true,
	DataFileMergeRatio: 0.5,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 10000,
	SyncWrites:  true,
}
