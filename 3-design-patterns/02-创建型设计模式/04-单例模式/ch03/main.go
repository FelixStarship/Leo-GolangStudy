package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例模式：并发安全
// 通过原子操作保证单例对象并发安全
// mem model:内存模型
var initialized uint32
var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	// 内存标记
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if initialized == 0 {
		instance = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象，格子衫来")
}

func main() {

	s := GetInstance()
	s.SomeThing()

}
