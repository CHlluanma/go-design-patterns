package singleton_pattern

import "sync"

// doubleC 在懒汉式的基础上优化，减少加锁操作，保证在并发安全下不影响性能
type doubleC struct {
	value int
}

// 声明私有变量
var doubleCInstance *doubleC

// 加锁，并发安全
var mutexC sync.Mutex

// GetDoubleCInstance 双重检查，同时并发安全
func GetDoubleCInstance() *doubleC {
	if doubleCInstance == nil {
		mutexC.Lock()
		if doubleCInstance == nil {
			doubleCInstance = new(doubleC)
		}
		mutex.Unlock()
	}
	return doubleCInstance
}
