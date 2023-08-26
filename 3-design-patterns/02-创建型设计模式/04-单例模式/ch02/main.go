package main

import (
	"fmt"
	"sync"
)

// 单例模式：懒汉式

// 保证并发安全
var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(singleton)
		return instance
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
