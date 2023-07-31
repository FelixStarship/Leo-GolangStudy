package main

import "fmt"

// 1.简单工厂

// ==============抽象层================//
// 水果类
type Fruit interface {
	Show()
}

type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("我是梨")
}

// ================工厂模块===================== //
type Factory struct{}

func (f *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

func main() {

	// ==========业务逻辑层==============
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()

}
