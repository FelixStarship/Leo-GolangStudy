package main

import "fmt"

type Book struct {
	pages int
}

type Books []Book

// 指针运算【共享间接部分】
func (books *Books) Modify() {
	(*books)[0].pages = 500
	*books = append(*books, Book{789})
}

func main() {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books)
}
