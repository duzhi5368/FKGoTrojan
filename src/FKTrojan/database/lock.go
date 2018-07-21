package database

import "sync"

var (
	// 避免竞争 这里简单做lock 同一时间只有一个goroutine可以操作database
	// 服务端无压力 不存在性能问题
	db_lock sync.Mutex
)

func Lock() {
	db_lock.Lock()
}
func Unlock() {
	db_lock.Unlock()
}
