package main

import "fmt"

// 1.依赖倒转原则
type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

type BMW struct {
}

func (m *BMW) Run() {
	fmt.Println("BMW is running....")
}

type Zhang3 struct {
}

func (z *Zhang3) DriveBenZ(benz *Benz) {
	fmt.Println("zhang3 Drive Benz")
	benz.Run()
}

func (z *Zhang3) DriveBMW(bmw *BMW) {
	fmt.Println("zhang3 Drive BMW")
	bmw.Run()
}

type Li4 struct {
}

func (l *Li4) DriveBenZ(ben *Benz) {
	fmt.Println("li4 Drive Benz")
	ben.Run()
}

func (l *Li4) DriveBMW(bmw *BMW) {
	fmt.Println("li4 drive BMW")
	bmw.Run()
}

/*


func main() {
	//业务1 张3开奔驰
	benz := &Benz{}
	zhang3 := &Zhang3{}
	zhang3.DriveBenZ(benz)

	//业务2 李四开宝马
	bmw := &BMW{}
	li4 := &Li4{}
	li4.DriveBMW(bmw)
}
*/

func main() {
	benz := &Benz{}
	zhang3 := Zhang3{}
	zhang3.DriveBenZ(benz)

	bmw := &BMW{}
	li4 := Li4{}
	li4.DriveBMW(bmw)
}
