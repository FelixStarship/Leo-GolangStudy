package main

import "fmt"

// 1.容器【元素】的可寻址性

/*a := [5]int{2, 3, 5, 7}
s := make([]bool, 2)
pa2, ps1 := &a[2], &s[1]
fmt.Println(*pa2, *ps1) // 5 false
a[2], s[1] = 99, true
fmt.Println(*pa2, *ps1) // 99 true
ps0 := &[]string{"Go", "C"}[0]
fmt.Println(*ps0) // Go

m := map[int]bool{1: true}
_ = m
// 下面这几行编译不通过。
/*
	_ = &[3]int{2, 3, 5}[0]
	_ = &map[int]bool{1: true}[1]
	_ = &m[1]
*/
// -----------------------------------------------
/*
 a := [5]int{2, 3, 5, 7}
	s := make([]bool, 2)
	pa2, ps1 := &a[2], &s[1] // 寻址
	fmt.Println(*pa2, *ps1)  // 解引用

	a[2], s[1] = 99, true
	fmt.Println(*pa2, *ps1) // 99 true

	ps0 := &[]string{"gzsl", "Go"}
	fmt.Println(*ps0)

	m := map[int]bool{1: true} // 所有的映射元素都不可以寻址

	fmt.Println(&m)
	fmt.Println(&a[1])
*/

/*
  type T struct{age int}
	mt := map[string]T{}
	mt["John"] = T{age: 29} // 整体修改是允许的
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789} // 整体修改是允许的

	// 这两个赋值编译不通过，因为部分修改一个映射
	// 元素是非法的。这看上去确实有些反直觉。
	/*
	ma[1][1] = 123      // error
	mt["John"].age = 30 // error
*/

// 读取映射元素的元素或者字段是没问题的。
// fmt.Println(ma[1][1])       // 789
// fmt.Println(mt["John"].age) // 29

// ---------------------------------------------------

/*
 type T struct {
		age int
	}

	mt := map[string]T{}
	mt["gzsl"] = T{age: 300}

	// mt["gzsl"].age=600 不支持部分修改

	fmt.Println(mt["gzsl"])

	ma := map[int][5]int{}
	ma[1] = [5]int{789, 567}

	//ma[1][0]=900  不支持部分修改

	//---------------------------------------------

	// 引入中间变量
	t := mt["gzsl"]
	t.age = 100

	mt["gzsl"] = t

	fmt.Println(mt)

	a := ma[1]
	a[0] = 100
	ma[1] = a
	fmt.Println(ma)
*/

// 1.从【数组】或者【切片】派生切片（取子切片）

/*
 a := make([]int, 3, 4)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a, len(a), cap(a))

	a1 := a[:4]                       // start:end
	fmt.Println(a1, len(a1), cap(a1)) // [1 2 3 0] 4 4

	a1[0] = 100

	fmt.Println(a, a1)
*/

// ------------------------------------------------------

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s0 := a[:]     // <=> s0 := a[0:7:7]
	s1 := s0[:]    // <=> s1 := s0
	s2 := s1[1:3]  // <=> s2 := a[1:3]
	s3 := s1[3:]   // <=> s3 := s1[3:7]
	s4 := s0[3:5]  // <=> s4 := s0[3:5:7]
	s5 := s4[:2:2] // <=> s5 := s0[3:5:5]
	s6 := append(s4, 77)
	s7 := append(s5, 88)
	s8 := append(s7, 66)
	s3[1] = 99
	fmt.Println(len(s2), cap(s2), s2) // 2 6 [1 2]
	fmt.Println(len(s3), cap(s3), s3) // 4 4 [3 99 77 6]
	fmt.Println(len(s4), cap(s4), s4) // 2 4 [3 99]
	fmt.Println(len(s5), cap(s5), s5) // 2 2 [3 99]
	s6[0] = 10000
	fmt.Println(len(s6), cap(s6), s6) // 3 4 [3 99 77]
	fmt.Println(len(s5), cap(s5), s5)

	fmt.Println(len(s7), cap(s7), s7) // 3 4 [3 4 88]
	fmt.Println(len(s8), cap(s8), s8) // 4 4 [3 4 88 66]

}
