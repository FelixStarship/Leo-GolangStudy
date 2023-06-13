package main

import "fmt"

// 2.结构体字段的可【寻址性】
//package main
//
//import "fmt"
//
//func main() {
//	type Book struct {
//		Pages int
//	}
//	var book = Book{} // 变量值book是可寻址的
//	p := &book.Pages
//	*p = 123
//	fmt.Println(book) // {123}
//
//	// 下面这两行编译不通过，因为Book{}是不可寻址的，
//	// 继而Book{}.Pages也是不可寻址的。
//	/*
//		Book{}.Pages = 123
//		p = &Book{}.Pages // <=> p = &(Book{}.Pages)
//	*/
//}

//type Book struct {
//	Pages int
//}
//
//var book =Book{}
//// .优先级高于&优先级
//p:=&book.Pages
//*p=123
//fmt.Println(book)
//
//
//a:=Book{}.Pages
//fmt.Println(a)
//
////// 组合字面量不可寻址
////Book{}.Pages=123
////// 编译不通过
////p2:=&(Book{}.Pages)
//
//// 隐式解引用
//p3:=&Book{Pages: 100}
//(*p3).Pages=200

// 3.隐式解引用

/*
type Book struct {
		pages int
	}
	book1 := &Book{100} // book1是一个指针
	book2 := new(Book)  // book2是另外一个指针
	// 像使用结构值一样来使用结构体值的指针。
	book2.pages = book1.pages
	// 上一行等价于下一行。换句话说，上一行
	// 两个选择器中的指针属主将被自动解引用。
	(*book2).pages = (*book1).pages
*/

// 4.结构体值类型转换

/*
package main

type S0 struct {
	y int "foo"
	x bool
}

type S1 = struct { // S1是一个无名类型
	x int "foo"
	y bool
}

type S2 = struct { // S2也是一个无名类型
	x int "bar"
	y bool
}

type S3 S2 // S3是一个定义类型（因而具名）。
type S4 S3 // S4是一个定义类型（因而具名）。
// 如果不考虑字段标签，S3（S4）和S1的底层类型一样。
// 如果考虑字段标签，S3（S4）和S1的底层类型不一样。

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}
func f() {
	v1 = S1(v2); v2 = S2(v1)
	v1 = S1(v3); v3 = S3(v1)
	v1 = S1(v4); v4 = S4(v1)
	v2 = v3; v3 = v2 // 这两个转换可以是隐式的
	v2 = v4; v4 = v2 // 这两个转换也可以是隐式的
	v3 = S3(v4); v4 = S4(v3)
}
*/

// 5.匿名结构体声明
/*
var aBook = struct {
	author struct { // 此字段的类型为一个匿名结构体类型
		firstName, lastName string
		gender              bool
	}
	title string
	pages int
}{
	author: struct {
		firstName, lastName string
		gender              bool
	}{
		firstName: "Mark",
		lastName: "Twain",
	}, // 此组合字面量中的类型为一个匿名结构体类型
	title: "The Million Pound Note",
	pages: 96,
}
*/
func main() {

	var aBook = struct {
		author struct {
			firstName, lastName string
			gender              bool
		}
		title string
		pages int
	}{
		author: struct {
			firstName, lastName string
			gender              bool
		}{firstName: "格子衫来", lastName: "格子衫来", gender: true},
		title: "格子衫来",
		pages: 101,
	}

	fmt.Println(aBook)
}
