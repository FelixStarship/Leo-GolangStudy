package main

import "fmt"

// ===== >   实现层  < ========
type Car interface {
	Run()
}

type Drive interface {
	Drive(car Car)
}

// ===== >   实现层  < ========
type BenZ struct{}

func (b *BenZ) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct{}

func (m *Bmw) Run() {
	fmt.Println("BMW is running...")
}

type Zhang3 struct{}

func (z *Zhang3) Drive(car Car) {
	fmt.Println("zhang3 drive car")
	car.Run()
}

type Li4 struct{}

func (l *Li4) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}

func main() {

	// ===== >   业务逻辑层  < ========

	// 张三 开宝马
	var bmw Car
	bmw = &Bmw{}

	var zhang3 Drive
	zhang3 = &Zhang3{}
	zhang3.Drive(bmw)

	// 李四 开奔驰
	var benz Car
	benz = &BenZ{}

	var li4 Drive
	li4 = &Li4{}
	li4.Drive(benz)

}
