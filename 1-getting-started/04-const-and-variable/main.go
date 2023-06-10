package main

// 1.常量声明
const Π = 3.1416

const Pi = Π // 赋值操作

// 一组具名常量
const (
	No         = !Yes
	Yes        = true
	MaxDegress = 360
	Uint       = "弧度"
)

// 类型确定具名常量

const X float32 = 3.14

//const X =float32(3.14)

const (
	A, B int64   = -3, 5
	Y    float32 = 2.1718
)

/*const (
	A, B = int64(-3), int64(5)
	Y    = float32(2.718)
)*/

// 2.常量声明自动补全
//const (
//	x float32 = 3.14
//	y
//	z
//	a, b = "格子衫来", "格子衫来"
//	c, _ //(【_】空标识符)
//)

// 自动补全
const (
	x float32 = 3.14
	y float32 = 3.14
	z float32 = 3.14

	a, b = "格子衫来", "格子衫来"
	c, _ = "格子衫来", "格子衫来"
)

// 3.iota
const (
	k = 3 //iota==0  ==> 0+3

	m float32 = iota + .5 // iota+.5 ==>1+.5

	n //自动补全:iota+.5 ===> 2+.5
)

// 实际编程中
const (
	Failed    = iota - 1 // =-1  0-1
	Unknown              // iota-1 1 1-1
	Succeeded            // iota-1 ==> 2-1
)

// 4.变量声明和赋值操作语句
// 4.1标准变量声明形式
/*
  var 关键字 ：var [标识符](变量名)[数据类型]
*/

var lang, website string = "格子衫来", "https://gezishlai.com"
var compiled, daynamic bool = true, false
var announceYear int = 2023

// ***和常量声明一样，多个类型的【变量】可以声明在一条语句中***

/*var lang, dynamic = "格子衫来", false

var website = "https://gezishlai.com"
*/

// 多个变量用一对小括号组团在一起
var (
	study, bornYear, compileds     = "格子衫来", 2023, true
	announceAt, releaseAt      int = 2023, 2024
	createBy, web              string
)

// 5.变量和常量的作用域
/*
{
    x,y:=x,y-10   // x和y遮挡了外层声明的x和y
}
*/

func main() {
	// 声明三个局部常量
	const DoublePi, HalfPi, Unit2 = Π * 2, Π * 0.5, "度"

	// 4.2短变量的声明形式
	lang, year := "格子衫来", 2007

	year, createBy := 2009, "格子衫来 Research"

	// 这是一个纯赋值语句
	lang, year = "格子衫来", 2023

	print(lang, "由", createBy, "发明")

	print("并发布于", year, "年")

	println()
}
