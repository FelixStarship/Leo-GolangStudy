package main

import "fmt"

// 在go语言中为什么需要指针？？

/*


package main

import "fmt"

 // 因为在Go中，所有的赋值（包括函数调用传参）过程都是一个值【复制】过程。 所以在上面的double函数体内修改的是变量a的一个【副本】，而没有修改变量a本身。
func double(x int) {
	x += x
}

func main() {
	var a = 3
	double(a)
	fmt.Println(a) // 3
}
*/

/*
package main

import "fmt"

	func double(x *int) {
		*x += *x
		x = nil // 此行仅为讲解目的
	}

func main() {

		var a = 3
		double(&a)
		fmt.Println(a) // 6

	    // 在此函数体内对传入的指针实参的修改x = nil依旧不能反映到函数外，因为此修改发生在此指针的一个副本上。 所以在double函数调用之后，局部变量p的值并没有被修改为nil。
		p := &a
		double(p)
		fmt.Println(a, p == nil) // 12 false
	}
*/
func double(x *int) {
	*x += *x
	x = nil

	fmt.Println(&x)

}

// 在Go中返回一个局部变量的地址是安全的
func newInt() *int {
	a := 3
	return &a
}

//package main
//
//import "fmt"
//
//func main() {
//	a := int64(5)
//	p := &a
//
//	// 下面这两行编译不通过。
//	/*
//		p++
//		p = (&a) + 8
//	*/
//
//	*p++
//	fmt.Println(*p, a)   // 6 6
//	fmt.Println(p == &a) // true
//
//	*&a++
//	*&*&a++
//	**&p++
//	*&*p++
//	fmt.Println(*p, a) // 10 10
//}

func main() {

	//a := int64(5)
	//p:=&a

	// 编译不通过
	/*p++
	p=(&a)+8*/

}
