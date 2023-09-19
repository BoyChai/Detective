package blast

import "Detective/struct"

type Archiver interface {
	// Open 打开文件
	Open(string) _struct.Error
	// GetType 获取类型
	GetType() string
}
