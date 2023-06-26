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
func main() {

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
