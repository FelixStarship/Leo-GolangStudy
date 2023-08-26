package main

import "fmt"

// 单例模式：饿汉式

// 初始化私有
type singleton struct{}

// 全局变量，初始化实例
var instance *singleton = new(singleton)

// 定义实例的全局访问点
func GetInstance() *singleton {
	return instance
}

// 实例对应的成员方法
func (s *singleton) SomeThing() {
	fmt.Println("单例对象，格子衫来")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
