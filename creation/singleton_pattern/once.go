package singleton_pattern

import "sync"

type onceValue struct {
	value int
}

var onceInstance *onceValue

var once sync.Once

func GetOnceInstance() *onceValue {
	once.Do(func() {
		onceInstance = new(onceValue)
	})
	return onceInstance
}
