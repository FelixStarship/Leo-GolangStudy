package main

import "fmt"

// 1.切片转化为数组指针
/*
 package main

type S []int
type A [2]int
type P *A

func main() {
	var x []int
	var y = make([]int, 0)
	var x0 = (*[0]int)(x) // okay, x0 == nil
	var y0 = (*[0]int)(y) // okay, y0 != nil
	_, _ = x0, y0

	var z = make([]int, 3, 5)
	var _ = (*[3]int)(z) // okay
	var _ = (*[2]int)(z) // okay
	var _ = (*A)(z)      // okay
	var _ = P(z)         // okay

	var w = S(z)
	var _ = (*[3]int)(w) // okay
	var _ = (*[2]int)(w) // okay
	var _ = (*A)(w)      // okay
	var _ = P(w)         // okay

	var _ = (*[4]int)(z) // 会产生恐慌
}
*/
/*
 type S []int
type A [2]int
type P *A
// 没有被初始化切片 x==nil
	var x []int

	// make 声明 y！=nil
	var y = make([]int, 0)
	var x0 = (*[0]int)(x) //数组指针
	var y0 = (*[0]int)(y)
	fmt.Println(x0, y0)
	_, _ = x0, y0

	var z = make([]int, 3, 5)
	var _ = (*[3]int)(z)
	var _ = (*[2]int)(z)
	var _ = (*A)(z)
	var _ = P(z)

	var w = S(z)
	fmt.Println(w)
	var _ = (*[3]int)(w)
	var _ = (*[2]int)(w)
	var _ = (*A)(w)

	var _ = (*[4]int)(z) // 产生恐慌：panic: runtime error: cannot convert slice with length 3 to array or pointer to array with length 4

*/

// 1.使用内置copy函数来复制切片元素
/*
type Ta []int
	type Tb []int
	dest := Ta{1, 2, 3}
	src := Tb{5, 6, 7, 8, 9}
	n := copy(dest, src)
	fmt.Println(n, dest) // 3 [5 6 7]
	n = copy(dest[1:], dest)
	fmt.Println(n, dest) // 2 [5 5 6]

	a := [4]int{} // 一个数组
	n = copy(a[:], src)
	fmt.Println(n, a) // 4 [5 6 7 8]
	n = copy(a[:], a[2:])
	fmt.Println(n, a) // 2 [7 8 7 8]
*/
// ------------------------------------------------
/*
type Ta []int
	dest := Ta{1, 2, 3}
	src := Ta{5, 6, 7, 8, 9}
	n := copy(dest, src) //int copy函数返回复制了多少个元素，此值（int类型）为这两个切片的长度的较小值。
	fmt.Println(n, dest) // 3 [5 6 7]
	n = copy(dest[1:], dest)
	fmt.Println(n, dest)

	a := [4]int{}
	n = copy(a[:], src)
	fmt.Println(n, a)     // 4 [5 6 7 8]
	n = copy(a[:], a[2:]) // 2 [7 8 7 8]   此值（int类型）为这两个切片的长度的较小值。
*/

// 3.遍历容器元素
/*
 m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
	for lang, year := range m {
		fmt.Printf("%v: %v \n", lang, year)
	}

	a := [...]int{2, 3, 5, 7, 11}
	for i, prime := range a {
		fmt.Printf("%v: %v \n", i, prime)
	}

	s := []string{"go", "defer", "goto", "var"}
	for i, keyword := range s {
		fmt.Printf("%v: %v \n", i, keyword)
	}
*/
/*
// 映射
	m := map[string]int{"gzsl": 1996, "Go": 2009, "C++": 1983}
	for lang, year := range m {
		fmt.Printf("%v:%v \n", lang, year)
	}
	// 数组
	a := [...]int{2, 3, 5, 7, 11}
	for i, prime := range a {
		fmt.Printf("%v:%v \n", i, prime)
	}
	// 切片
	s := []string{"格子衫来", "defer", "goto", "var"}
	for i, keyword := range s {
		fmt.Printf("%v:%v \n", i, keyword)
	}
*/

func main() {
	// 遍历一个nil映射或者nil切片是允许的。这样的遍历可以看作是一个空操作。
	var a []int
	fmt.Println(a)
	for range a {

	}

	var a1 map[string]int
	fmt.Println(a1)
	for range a1 {
	}
}
