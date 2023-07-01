package main

import "fmt"

// 变长参数和变长参数函数类型

/*

 func (values ...int64) (sum int64)
 func (sep string, tokens ...string) string

*/

/*
// Sum返回所有输入实参的和。
func Sum(values ...int64) (sum int64) {
	// values的类型为[]int64。
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

// Concat是一个低效的字符串拼接函数。
func Concat(sep string, tokens ...string) string {
	// tokens的类型为[]string。
	r := ""
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return r
}
*/
/*
 a0 := Sum()
	a1 := Sum(2)
	a3 := Sum(2, 3, 5)
	// 上面三行和下面三行是等价的。
	b0 := Sum([]int64{}...) // <=> Sum(nil...)
	b1 := Sum([]int64{2}...)
	b3 := Sum([]int64{2, 3, 5}...)
	fmt.Println(a0, a1, a3) // 0 2 10
	fmt.Println(b0, b1, b3) // 0 2 10
*/

/*
	 tokens := []string{"Go", "C", "Rust"}
		langsA := Concat(",", tokens...)        // 风格1
		langsB := Concat(",", "Go", "C","Rust") // 风格2
		fmt.Println(langsA == langsB)           // true
*/

/*
a0 := Sum() //调用变长参数函数、传递0和或多个实参

	a1 := Sum(2)
	a3 := Sum(2, 3, 5)

	b0 := Sum([]int64{}...)
	b1 := Sum([]int64{2}...)
	b3 := Sum([]int64{2, 3, 5}...)

	fmt.Println(a0, a1, a3) // 0 2 10
	fmt.Println(b0, b1, b3) // 0 2 10

	tokens := []string{"gzsl", "格子衫来", "Rust"}
	langsA := Concat(",", tokens...)
	langsB := Concat(",", "gzsl", "格子衫来", "Rust")
	fmt.Println(langsA == langsB) //true
*/
func fa() int {
a:
	goto a // 死循环
}

// 2.函数值

/*
 func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n) // f的类型为"func(int) int"
}
*/

/*
 fmt.Printf("%T\n", Double) // func(int) int
	// Double = nil // error: Double是不可修改的

	var f func(n int) int // 默认值为nil
	f = Double
	g := Apply
	fmt.Printf("%T\n", g) // func(int, func(int) int) int

	fmt.Println(f(9))         // 18
	fmt.Println(g(6, Double)) // 12
	fmt.Println(Apply(6, f))  // 12
*/

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n)
}

/*
    fmt.Printf("%T\n", Double) // 函数类型 func(int) int
	//Double = nil // error: 函数值不可修改
	//fmt.Println(Double)
	var f func(n int) int

	fmt.Println(f)

	f = Double
	g := Apply //相同
	fmt.Printf("%T\n", g)

	fmt.Println(f(9))

	fmt.Println(g(6, Double))

	fmt.Println(Apply(6, f))
*/

// 3.匿名函数调用
/*
 // 此函数返回一个函数类型的结果，亦即闭包（closure）。
	isMultipleOfX := func (x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}

	var isMultipleOf3 = isMultipleOfX(3)
	var isMultipleOf5 = isMultipleOfX(5)
	fmt.Println(isMultipleOf3(6))  // true
	fmt.Println(isMultipleOf3(8))  // false
	fmt.Println(isMultipleOf5(10)) // true
	fmt.Println(isMultipleOf5(12)) // false

	isMultipleOf15 := func(n int) bool {
		return isMultipleOf3(n) && isMultipleOf5(n)
	}
	fmt.Println(isMultipleOf15(32)) // false
	fmt.Println(isMultipleOf15(60)) // true
*/
func main() {

	isMutipleOfX := func(x int) func(int) bool {
		// 闭包
		return func(n int) bool {
			return n%x == 0
		}
	}

	var isMutipleOf3 = isMutipleOfX(3)
	var isMutipleOf5 = isMutipleOfX(5)

	fmt.Println(isMutipleOf3(6))
	fmt.Println(isMutipleOf3(8))
	fmt.Println(isMutipleOf5(10))
	fmt.Println(isMutipleOf5(12))

}

func Sum(values ...int64) (sum int64) {
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

func Concat(sep string, tokens ...string) string {
	r := ""
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return r
}
