package main

import "fmt"

// 1.GO结构体声明

/*
package main

import (
	"fmt"
)

type Book struct {
	title, author string
	pages         int
}

func main() {
	book := Book{"Go语言101", "老貘", 256}
	fmt.Println(book) // {Go语言101 老貘 256}

	// 使用带字段名的组合字面量来表示结构体值。
	book = Book{author: "老貘", pages: 256, title: "Go语言101"}
	// title和author字段的值都为空字符串""，pages字段的值为0。
	book = Book{}
	// title字段空字符串""，pages字段为0。
	book = Book{author: "老貘"}

	// 使用选择器来访问和修改字段值。
	var book2 Book // <=> book2 := Book{}
	book2.author = "Tapir"
	book2.pages = 300
	fmt.Println(book2.pages) // 300
}
*/

type Book struct {
	title, author string
	pages         int
}

func main() {

	//1.
	book := Book{"GO语言从入门到精通", "格子衫来", 1000}
	fmt.Println(book)

	// 2.
	book1 := Book{title: "GO语言从入门到精通", author: "格子衫来", pages: 2000}
	fmt.Println(book1)

	boo3 := Book{}
	boo3 = Book{author: "格子衫来"}
	fmt.Println(boo3)

	var book4 Book
	book4.author = "格子衫来"
	book4.pages = 3000

	fmt.Println(book4)

}
