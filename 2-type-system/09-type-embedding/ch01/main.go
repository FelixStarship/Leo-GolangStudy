package main

import (
	"fmt"
	"reflect"
)

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

/*
 var gaga = Singer{Person: Person{"格子衫来", 18}}
	gaga.PrintName()

	gaga.Name = "gzsl"
	(&gaga).SetAge(31) // 指针类型（引用类型）
	(&gaga).PrintName()
	fmt.Println(gaga.Age)
*/

/*
 t := reflect.TypeOf(Singer{}) // the Singer type
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&Singer{}) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}
*/

func main() {

	t := reflect.TypeOf(Singer{})
	fmt.Println(t, "has", t.NumField(), "fields:")
	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(" field#", i, ":", t.Field(i).Name)
	}
	fmt.Println(t, "has", t.NumMethod(), "methods:")
	// 遍历方法
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(" method#", i, ":", t.Method(i).Name)
	}

	pt := reflect.TypeOf(&Singer{})
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println(" method#", i, ":", pt.Method(i).Name)
	}

}
