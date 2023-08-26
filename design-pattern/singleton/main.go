package main

import "sync"

var (
	once     sync.Once
	instance Singleton
)

type Singleton interface {
	GetName() string
}

type singleton struct {
	name string
}

func (s *singleton) GetName() string {
	return s.name
}

func GetInstance() Singleton {
	// 保证实例化方法只执行一次
	once.Do(func() {
		instance = &singleton{name: "singleton"}
	})
	return instance
}

func main() {
	println(GetInstance().GetName()) // singleton
}
