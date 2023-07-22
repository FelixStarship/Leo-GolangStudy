package main

import "fmt"

// 1.类型内嵌
/*
 1.1类型内嵌语法
 type P = *bool
	type M = map[int]int
	var x struct {
		string // 一个具名非指针类型
		error  // 一个具名接口类型
		*int   // 一个无名指针类型
		P      // 一个无名指针类型的别名
		M      // 一个无名类型的别名

		http.Header // 一个具名映射类型
	}
	x.string = "Go"
	x.error = nil
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}
*/

/*
 // 1.类型内嵌的语法
	type P = *bool
	type M = map[int]int
	var x struct {
		string //匿名字段<==>内嵌字段 类型名即为此字段的名称
		error
		*int
		P
		M

		http.Header
	}

	x.string = "Go"
	x.error = nil
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}
*/

/*
 1.2 类型内嵌的意义？
 type Person struct {
	Name string
	Age  int
}
func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person // 通过内嵌Person类型来扩展之
	works  []string
}
var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName() // Name: Gaga
	gaga.Name = "Lady Gaga"
	(&gaga).SetAge(31)
	(&gaga).PrintName()   // Name: Lady Gaga
	fmt.Println(gaga.Age) // 31
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person
	works []string
}

func main() {

	var gaga = Singer{Person: Person{"格子衫来", 18}}
	gaga.PrintName()

	gaga.Name = "gzsl"
	(&gaga).SetAge(31) // 指针类型（引用类型）
	(&gaga).PrintName()
	fmt.Println(gaga.Age)

}
