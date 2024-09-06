package fio

const DataFilePerm = 0644

// IOManager 抽象IO管理接口，可以接入不同的IO类型，目前支持标准文件IO
type IOManager interface {
	Read([]byte, int64) (int, error) // Read 从文件的给定位置读取对应数据
	Write([]byte) (int, error)       // Write 写入字节数组到文件中
	Sync() error                     // Sync 持久化数据
	Close() error                    // Close 关闭文件
	Size() (int64, error)            // Size 获取到文件大小
}

// NewIOManager 初始化 IOManager，目前只支持标准 FileIO
func NewIOManager(fileName string) (IOManager, error) {
	return NewFileIOManager(fileName)
}