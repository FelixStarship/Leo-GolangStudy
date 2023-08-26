package main

import "fmt"

// 1.工厂方法模式

// 抽象产品
type Fruit interface {
	Show()
}

// 抽象工厂
type AbstractFactory interface {
	CreateFruit() Fruit
}

// 具体产品
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

type JapanApple struct{}

func (ja *JapanApple) Show() {
	fmt.Println("我是日本苹果")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("我是香焦")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("我是梨")
}

// 实现工厂
type AppleFactory struct{}

func (fac *AppleFactory) CreateFruit() Fruit {

	var fruit Fruit

	fruit = new(Apple)

	return fruit
}

type JapanAppleFactory struct{}

func (fac *JapanAppleFactory) CreateFruit() Fruit {

	var fruit Fruit

	fruit = new(JapanApple)

	return fruit
}

type BananaFactory struct{}

func (fac *BananaFactory) CreateFruit() Fruit {

	var fruit Fruit

	fruit = new(Banana)

	return fruit
}

type PearFactory struct{}

func (fac *PearFactory) CreateFruit() Fruit {

	var fruit Fruit

	fruit = new(Pear)

	return fruit
}

func main() {

	// 负责生产苹果的工厂类
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)

	// 生产苹果
	var apple Fruit
	apple = appleFac.CreateFruit()

	apple.Show()

	// 负责生产香蕉工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)

	// 生产香蕉
	var banana Fruit
	banana = bananaFac.CreateFruit()
	banana.Show()

	// 负责生产梨的工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)

	// 生产梨
	var pear Fruit
	pear = pearFac.CreateFruit()
	pear.Show()

	// 负责生产日本苹果工厂
	var japanFac AbstractFactory
	japanFac = new(JapanAppleFactory)

	// 生产日本苹果
	var japanApple Fruit
	japanApple = japanFac.CreateFruit()
	japanApple.Show()

}
