package main

import "fmt"

//1.容器值的比较

//package main
//
//import "fmt"
//
//func main() {
//	var a [16]byte
//	var s []int
//	var m map[string]int
//
//	fmt.Println(a == a)   // true
//	fmt.Println(m == nil) // true
//	fmt.Println(s == nil) // true
//	fmt.Println(nil == map[string]int{}) // false
//	fmt.Println(nil == []int{})          // false
//
//	// 下面这些行编译不通过。
//	/*
//		_ = m == m
//		_ = s == s
//		_ = m == map[string]int(nil)
//		_ = s == []int(nil)
//		var x [16][]int
//		_ = x == x
//		var y [16]map[int]bool
//		_ = y == y
//	*/
//}
//---------------------------------
//var a [16]byte
//var s []int
//var m map[string]int
//
//// 数组类型可以比较
//fmt.Println(a == a)
//// map与零值比较
//fmt.Println(m == nil)
//// 切片与零值比较
//fmt.Println(s == nil)
//// 已经被初始化、非零值
//fmt.Println(nil == map[string]int{})
//fmt.Println(nil == []int{})
//
//fmt.Println(`任意两个映射值（或切片值）是不能相互比较的`)
//
///*
//	_=m==m
//	_=s==s
//	_=m==map[string]int(nil)
//	_=s==[]int(nil)
//
//	var x [16][]int
//	_=x==x*/

// 2.查看容器值的长度和容量
/*
package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(len(a), cap(a)) // 5 5
	var s []int
	fmt.Println(len(s), cap(s)) // 0 0
	s, s2 := []int{2, 3, 5}, []bool{}
	fmt.Println(len(s), cap(s), len(s2), cap(s2)) // 3 3 0 0
	var m map[int]bool
	fmt.Println(len(m)) // 0
	m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
	fmt.Println(len(m), len(m2)) // 2 0
}
*/
// -----------------------------------
/*
// 数组定长（数组的长度和容量无法改变）
	var a [5]int
	fmt.Println(len(a), cap(a)) // 5 //5

	b := [...]int{1, 2, 3, 5, 6, 7}
	fmt.Println(len(b), cap(b))

	var s []int
	fmt.Println(len(s), cap(s))

	s, s2 := []int{2, 3, 5}, []bool{}
	fmt.Println(len(s), cap(s), len(s2), len(s2))

	var m map[int]bool
	fmt.Println(len(m))

	m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
	fmt.Println(len(m), len(m2))
*/

// 3.读取和修改容器的元素
//a := [3]int{-1, 0, 1}
//s := []bool{true, false}
//m := map[string]int{"abc": 123, "xyz": 789}
//fmt.Println (a[2], s[1], m["abc"])    // 读取
//a[2], s[1], m["abc"] = 999, true, 567 // 修改
//fmt.Println (a[2], s[1], m["abc"])    // 读取
//
//n, present := m["hello"]
//fmt.Println(n, present, m["hello"]) // 0 false 0
//n, present = m["abc"]
//fmt.Println(n, present, m["abc"]) // 567 true 567
//m = nil
//fmt.Println(m["abc"]) // 0
//
//// 下面这两行编译不通过。
///*
//	_ = a[3]  // 下标越界
//	_ = s[-1] // 下标越界
//*/
//
//// 下面这几行每行都会造成一个恐慌。
//_ = a[n]         // panic: 下标越界
//_ = s[n]         // panic: 下标越界
//m["hello"] = 555 // panic: m为一个零值映射

func main() {

	// 索引从0开始
	a := [3]int{-1, 0, 1}
	s := []bool{true, false}
	m := map[string]int{"格子衫来": 123, "gzsl": 789}
	// 读取容器中的元素
	fmt.Println(a[2], s[1], m["gzsl"])
	// 修改容器元素
	a[2], s[1], m["gzsl"] = 999, true, 567
	fmt.Println(a[2], s[1], m["gzsl"])

	// 映射操作
	n, ok := m["hello"] // 不会keyerr 安全
	fmt.Println(n, ok, m["hello"])

	// 存在元素
	n, ok = m["gzsl"]
	fmt.Println(n, ok, m["gzsl"])

	// 切片索引无法推断、编译时刻、切片可以动态库扩容、只有在运行时刻估值
	// 产生恐慌 panic: runtime error: index out of range [10] with length 2
	fmt.Println(s[10])

	m = nil
	// 产生恐慌 m为零值映射 panic: assignment to entry in nil map
	m["hello"] = 789

}
