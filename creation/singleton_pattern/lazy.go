package singleton_pattern

import "sync"

// lazy 懒汉式：在创建对象时不直接创建对象，在加载配置文件时才创建对象
type lazy struct {
	value int
}

// 声明私有变量
var lazyInstance *lazy

// 加锁，并发安全
var mutex sync.Mutex

// GetLazyInstance 获取单例对象
func GetLazyInstance() *lazy {
	mutex.Lock()
	defer mutex.Unlock()
	if lazyInstance == nil {
		lazyInstance = new(lazy)
	}
	return lazyInstance
}
