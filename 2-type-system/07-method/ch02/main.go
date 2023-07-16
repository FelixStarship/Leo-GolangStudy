package main

import "fmt"

// 2.属主参数的传参是值复制过程

/*
 type Book struct {
	pages int
}

func (b Book) SetPages(pages int) {
	b.pages = pages
}
*/

type Book struct {
	pages int
}

func (b Book) SetPages(pages int) {
	b.pages = pages
}

type Books []Book

func (books Books) Modify() {

	// 属主参数直接部分修改(新的切片)  切片扩容
	books = append(books, Book{100})
	// 属主参数间接部分修改
	books[0].pages = 500

	fmt.Println(len(books), cap(books)) // 3 4

}

/*
    var b Book
	b.SetPages(123)
	fmt.Println(b.pages)   方法体内的修改不会反应到方法体外
*/

func main() {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books)
	fmt.Println(len(books), cap(books)) // 2 2
}
