package main

import "fmt"

// 1.方法声明

/*
 // Age和int是两个不同的类型。我们不能为int和*int
// 类型声明方法，但是可以为Age和*Age类型声明方法。
type Age int
func (age Age) LargerThan(a Age) bool {
	return age > a
}
func (age *Age) Increase() {
	*age++
}

// 为自定义的函数类型FilterFunc声明方法。
type FilterFunc func(in int) bool
func (ff FilterFunc) Filte(in int) bool {
	return ff(in)
}

// 为自定义的映射类型StringSet声明方法。
type StringSet map[string]struct{}
func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// 为自定义的结构体类型Book和它的指针类型*Book声明方法。

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

*/

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}

func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

type FilterFunc func(in int) bool

func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

type Age int

// 值类型属主
func (age Age) LargerThan(a Age) bool {
	return age > a
}

// age 指针类型属主
func (age *Age) Increase() {
	*age++
}

// 2.每个方法对应着一个隐式声明的函数
/*
func Book.Pages(b Book) int {
	return b.pages // 此函数体和Book类型的Pages方法体一样
}

func (*Book).SetPages(b *Book, pages int) {
	b.pages = pages // 此函数体和*Book类型的SetPages方法体一样
}
*/

//func Book.Pages(b Book) int{
//
//}

func main() {

	// 调用方法隐式声明的的【函数】
	var book Book
	book.pages = 100
	(*Book).SetPages(&book, 100)
	fmt.Println(Book.Pages(book))
}
