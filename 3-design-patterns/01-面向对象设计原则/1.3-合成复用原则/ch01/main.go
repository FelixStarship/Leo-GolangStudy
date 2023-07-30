package main

import "fmt"

// 3.合成复用原则

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

type CatC struct {
	C *Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

/*


func main() {
	//通过继承增加的功能，可以正常使用
	cb := new(CatB)
	cb.Eat()
	cb.Sleep()

	//通过组合增加的功能，可以正常使用
	cc := new(CatC)
	cc.C = new(Cat)
	cc.C.Eat()
	cc.Sleep()
}
*/

func main() {
	cb := new(CatB)
	cb.Eat()
	cb.Sleep()

	cc := new(CatC)
	cc.C.Eat()
	cc.Sleep()
}
