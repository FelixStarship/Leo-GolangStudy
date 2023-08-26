package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) SomeThing() {

	fmt.Println("单例对象，格子衫来")
}

func main() {

	s := GetInstance()
	s.SomeThing()

}
