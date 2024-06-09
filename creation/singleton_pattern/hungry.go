package singleton_pattern

// hungry 直接通过初始化函数创建对象，是并发安全的
type hungry struct {
	value int
}

var hungryInstance *hungry

func init() {
	if hungryInstance == nil {
		hungryInstance = new(hungry)
	}
}

func GetHungryInstance() *hungry {
	return hungryInstance
}
