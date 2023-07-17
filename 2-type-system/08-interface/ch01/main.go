package main

import "fmt"

// 1.接口类型定义

/*
 // 此接口直接指定了两个方法和内嵌了两个其它接口。
// 其中一个为类型名称，另一个为类型字面表示形式。
type ReadWriteCloser = interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	error                      // 一个类型名称
	interface{ Close() error } // 一个类型字面表示形式
}

// 此接口类型内嵌了一个近似类型。
// 它的类型集由所有底层类型为[]byte的类型组成。
type AnyByteSlice = interface {
	~[]byte
}

// 此接口类型内嵌了一个类型并集。它的类型集包含6个类型：
// uint、uint8、uint16、uint32、uint64和uintptr。
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}
*/

type ReadWriteCloser = interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	error
	interface{ Close() error }
}

type AnyByteSlice = interface {
	~[]byte
}

// 2.实现接口
/*
type Aboutable interface {
	About() string
}

type Book struct {
	name string
	// 更多其它字段……
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}
*/

/*
 // 一个*Book值被包裹在了一个Aboutable值中。
	var a Aboutable = &Book{"Go语言101"}
	fmt.Println(a) // &{Go语言101}

	// i是一个空接口值。类型*Book实现了任何空接口类型。
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable实现了空接口类型interface{}。
	i = a
	fmt.Println(i) // &{Go语言101}
*/
// 隐式实现【关系】
type Aboutable interface {
	About() string
}

type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book:" + book.name
}

type MyInt int

func (MyInt) About() string {
	return "格子衫来"
}

func main() {

	var a Aboutable = &Book{"格子衫来"}
	fmt.Println(a)

	var i interface{} = &Book{"gzsl"} // 空接口==>Any
	fmt.Println(i)

	i = a // go语言中任何类型都实现了空接口
	fmt.Println(i)

}
