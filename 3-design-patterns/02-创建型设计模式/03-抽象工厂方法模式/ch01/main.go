package main

import "fmt"

// 1.抽象工厂方法模式

// 抽象产品
type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}
type AbstractPear interface {
	ShowPear()
}

// 抽象工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("中国梨")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {

	var apple AbstractApple

	apple = new(ChinaApple)

	return apple

}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {

	var banana AbstractBanana

	banana = new(ChinaBanana)

	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {

	var pear AbstractPear

	pear = new(ChinaPear)

	return pear
}

type JapanApple struct{}

func (ja *JapanApple) ShowApple() {
	fmt.Println("日本苹果")
}

type JapanBanana struct{}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("日本香蕉")
}

type JapanPear struct{}

func (jp *JapanPear) ShowPear() {
	fmt.Println("日本梨")
}

type JapanFactory struct{}

func (jf *JapanFactory) CreateApple() AbstractApple {

	var apple AbstractApple
	apple = new(JapanApple)

	return apple

}

func (jf *JapanFactory) CreateBanana() AbstractBanana {

	var banan AbstractBanana
	banan = new(JapanBanana)
	return banan
}

func (jf *JapanFactory) CreatePear() AbstractPear {

	var pear AbstractPear

	pear = new(JapanPear)

	return pear
}

type AmericanApple struct{}

func (aa *AmericanApple) ShowApple() {
	fmt.Println("美国苹果")
}

type AmericanBanana struct{}

func (ab *AmericanBanana) ShowBanana() {
	fmt.Println("美国香蕉")
}

type AmericanPear struct{}

func (ap *AmericanPear) ShowPear() {
	fmt.Println("美国梨")
}

type AmericanFactory struct{}

func (af *AmericanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(AmericanApple)
	return apple
}

func (af *AmericanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(AmericanBanana)
	return banana
}

func (af *AmericanFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(AmericanPear)
	return pear
}

func main() {

	// 1.初始化美国工厂
	var aFac AbstractFactory
	aFac = new(AmericanFactory)

	var aApple AbstractApple
	aApple = aFac.CreateApple()
	aApple.ShowApple()

	var aBanana AbstractBanana
	aBanana = aFac.CreateBanana()
	aBanana.ShowBanana()

	var aPear AbstractPear
	aPear = aFac.CreatePear()
	aPear.ShowPear()

	// 2.初始化中国工厂
	var cFac AbstractFactory
	cFac = new(ChinaFactory)

	var cApple AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()

	var cBanana AbstractBanana
	cBanana = cFac.CreateBanana()
	cBanana.ShowBanana()

	var cPear AbstractPear
	cPear = cFac.CreatePear()
	cPear.ShowPear()
}

/*
 练习：
	设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
	现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
	要求组装两台电脑，
			    1台（Intel的CPU，Intel的显卡，Intel的内存）
			    1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	用抽象工厂模式实现。
*/
