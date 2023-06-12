package main

import (
	"fmt"
)

/*
 指针
 *int  // 一个基类型为int的无名指针类型。
**int // 一个多级无名指针类型，它的基类型为*int。

type Ptr *int // Ptr是一个具名指针类型，它的基类型为int。
type PP *Ptr  // PP是一个具名多级指针类型，它的基类型为Ptr。
*/

/*
 1.声明指针(寻址)
   - new()
   - &
 2.指针取值操作（解引用）
   *T
*/

/*
    p0 := new(int)   // p0指向一个int类型的零值
	fmt.Println(p0)  // （打印出一个十六进制形式的地址）
	fmt.Println(*p0) // 0

	x := *p0              // **x是p0所引用的值的一个复制**。
	p1, p2 := &x, &x      // p1和p2中都存储着x的地址。
	                      // x、*p1和*p2表示着同一个int值。
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false

-------------------------------------------------------------------------------------
	p3 := &*p0            // <=> p3 := &(*p0)
	                      // <=> p3 := p0
	                      // p3和p0中存储的地址是一样的。
	fmt.Println(p0 == p3) // true
------------------------------------赋值操作-------------------------------------------------
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
*/

func main() {

	p0 := new(int)
	fmt.Println(p0)  //<=> p0所指向的int值的内存地址
	fmt.Println(&p0) // <==> *p0指针本身的内存地址
	fmt.Println(*p0) // <==> 值

	x := *p0

	p1, p2 := &x, &x
	fmt.Println(p1 == p2)
	fmt.Println(p0 == p1)

	p3 := &*p0
	fmt.Println(p0 == p3)

	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3)

	fmt.Printf("%T,%T \n", *p0, x)
	fmt.Printf("%T,%T \n", p0, p1)

}
