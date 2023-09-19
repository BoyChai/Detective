package blast

type Archiver interface {
	// Open 打开文件
	Open(string)
	// GetType 获取类型
	GetType() string
}
