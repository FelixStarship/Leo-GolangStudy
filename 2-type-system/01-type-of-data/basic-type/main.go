package main

import "fmt"

/*
  精通golang语言必学

概念：基本类型（basic type）
    内置字符串类型：string.
    内置布尔类型：bool.
    内置数值类型：
        int8、uint8（byte）、int16、uint16、int32（rune）、uint32、int64、uint64、int、uint、uintptr。
        float32、float64。
        complex64、complex128。
-----------------------------------------------------------------------------------------------------
概念：组合类型（composite type）

    指针类型 - 类C指针
    结构体类型 - 类C结构体
    函数类型 - 函数类型在Go中是一种一等公民类别
    容器类型，包括:
        数组类型 - 定长容器类型
        切片类型 - 动态长度和容量容器类型
        映射类型（map）- 也常称为字典类型。在标准编译器中映射是使用哈希表实现的。
    通道类型 - 通道用来同步并发的协程
    接口类型 - 接口在反射和多态中发挥着重要角色

*/

/*
// 假设T为任意一个类型，Tkey为一个支持比较的类型。

*T         // 一个指针类型
[5]T       // 一个元素类型为T、元素个数为5的数组类型
[]T        // 一个元素类型为T的切片类型
map[Tkey]T // 一个键值类型为Tkey、元素类型为T的映射类型

// 一个结构体类型

	struct {
		name string
		age  int
	}

// 一个函数类型
func(int) (bool, string)

// 一个接口类型

	interface {
		Method0(string) int
		Method1() (int, bool)
	}

// 几个通道类型
chan T
chan<- T
<-chan T
*/

/*
  语法：类型定义（type definition declaration）
  // 定义单个类型。
type NewTypeName SourceType

// 定义多个类型（将多个类型描述合并在一个声明中）。
type (
	NewTypeName1 SourceType1
	NewTypeName2 SourceType2
)

*/

// 1.一个新定义的类型和它的源类型为两个不同的类型。
// 2.在两个不同的类型定义中所定义的两个类型肯定是两个不同的类型。
// 3.一个新定义的类型和它的源类型的底层类型（将在下面介绍）一致并且它们的值可以相互显式转换。

/*
	   i := 0
		j := NewTypeName(10)

		i1 := NewTypeName1(1)
		i2 := NewTypeName2(2)

		if i == j {

		}

		if int(i1) == int(i2) {

		}
*/
type NewTypeName int

type (
	NewTypeName1 NewTypeName
	NewTypeName2 NewTypeName
)

/*
 // 下面这些新定义的类型和它们的源类型都是基本类型。
// 它们的源类型均为预声明类型。
type (
	MyInt int
	Age   int
	Text  string
)

// 下面这些新定义的类型和它们的源类型都是组合类型。
// 它们的源类型均为无名类型（见下下节）。
type IntPtr *int
type Book struct{author, title string; pages int}
type Convert func(in0 int, in1 bool)(out0 int, out1 string)
type StringArray [5]string
type StringSlice []string

func f() {
	// 这三个新定义的类型名称只能在此函数内使用。
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{Read([]byte) int}
}

*/

type (
	MyInt int
	Age   int
	Text  string
)

type IntPtr *int
type Book struct {
	author, title string
	pages         int
}

type StringArray [5]string
type StringSlice []string

/*
    // 自定义类型定义可以出现在函数体内。
	// map类型
	type PersonAge map[string]int
	// 通道类型可读、可写
	type MessageQueue chan string
	// 接口类型
	type Reader interface {
		Read([]byte) int
	}

*/

// 类型别名声明（type alias declaration）

/*
 type (
	Name = string
	Age  = int
)

type table = map[string]int
type Table = map[Name]Age

*/

func main() {

	// 类型别名申明语法
	type (
		Name = string
		Age  = int
	)

	type table = map[string]int
	type Table = map[Name]Age

	type (
		test string
	)

	// 初始化变量
	tb := table{
		"格子衫来":    0,
		"格子衫来...": 1,
	}

	fmt.Println(tb["格子衫来"])
}
