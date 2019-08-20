package memory_conn

import "sync"

var (
	MemoryDb *sync.Map
)

func init() {
	// 初始化内存数据库
	MemoryDb = &sync.Map{}
}
