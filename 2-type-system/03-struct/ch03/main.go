package main

import "fmt"

// 3.结构体内存分布

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "格子衫来"

	var p2 = &p1

	fmt.Println(p2.Age)
	p2.Name = "格子衫来，格子衫来"

	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)
	fmt.Printf("p1的地址%p \n", &p1)
	fmt.Printf("p2的地址%p  p2的值%p \n", &p2, p2)
}
