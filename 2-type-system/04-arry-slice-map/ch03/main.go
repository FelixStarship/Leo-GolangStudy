package main

import "fmt"

// 1.容器赋值
/*
   package main

import "fmt"

func main() {
	m0 := map[int]int{0:7, 1:8, 2:9}
	m1 := m0
	m1[0] = 2
	fmt.Println(m0, m1) // map[0:2 1:8 2:9] map[0:2 1:8 2:9]

	s0 := []int{7, 8, 9}
	s1 := s0
	s1[0] = 2
	fmt.Println(s0, s1) // [2 8 9] [2 8 9]

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println(a0, a1) // [7 8 9] [2 8 9]
}

*/
// ---------------------------------------------------------------------
/*
 // 初始化映射
	m0 := map[int]int{0: 7, 1: 8, 2: 9}
	m1 := m0 //共享底层元素
	m1[0] = 2

	fmt.Println(m0, m1)

	s0 := []int{7, 8, 9}
	s1 := s0 //共享底层元素
	s1[0] = 2
	fmt.Println(s0, s1)

	a0 := [...]int{7, 8, 9}
	a1 := a0 //值拷贝 复制过程
	a1[0] = 2
	fmt.Println(a0, a1)
*/

// 2.添加和删除容器元素
/*
 m := map[string]int{"Go": 2007}
	m["C"] = 1972     // 添加
	m["Java"] = 1995  // 添加
	fmt.Println(m)    // map[C:1972 Go:2007 Java:1995]
	m["Go"] = 2009    // 修改
	delete(m, "Java") // 删除
	fmt.Println(m)    // map[C:1972 Go:2009]
*/
//---------------------------------------
/*
 m := map[string]int{"格子衫来": 1996}
	m["c"] = 1972  //添加
	m["go"] = 2007 //添加
	fmt.Println(m)
	m["go"] = 2009 //修改

	delete(m, "c") //删除

	fmt.Println(m)

	delete(m, "gzsl") //空操作、不会产生恐慌
*/

/*
	  s0 := []int{2, 3, 5}
		fmt.Println(s0, cap(s0)) // [2 3 5] 3
		s1 := append(s0, 7)      // 添加一个元素
		fmt.Println(s1, cap(s1)) // [2 3 5 7] 6


		s2 := append(s1, 11, 13) // 添加两个元素
		fmt.Println(s2, cap(s2)) // [2 3 5 7 11 13] 6
		s3 := append(s0)         // <=> s3 := s0
		fmt.Println(s3, cap(s3)) // [2 3 5] 3
		s4 := append(s0, s0...)  // 以s0为基础添加s0中所有的元素
		fmt.Println(s4, cap(s4)) // [2 3 5 2 3 5] 6

		s0[0], s1[0] = 99, 789
		fmt.Println(s2[0], s3[0], s4[0]) // 789 99 2
*/

// ----------------------------------------------------------------
// 切片扩容

/*

	s0 := []int{2, 3, 5}
	fmt.Println(s0, cap(s0)) // [2 3 5] 3

	s1 := append(s0, 7) // s1编译器开辟新的内存产生s1切片（s0没有足够的[槽位]）

	fmt.Println(s1, cap(s1)) // [2 3 57] 6

	s1[0] = 0 //修改s1切片索引下标为0的元素《==》0

	fmt.Println(s0, s1) // [2 3 5] [0 3 5 7] s0和s1不共享元素、

	s2 := append(s1, 11, 13)
	fmt.Println(s2, cap(s2)) // [0 3 5 7 11 13] 6 s2<=>s1共享元素,编译器没有扩容，产生新的内存地址

	s3 := append(s0)         // s0<=>s3
	fmt.Println(s3, cap(s3)) // [2 3 5] 3 s3和s0共享元素

	s4 := append(s0, s0...)
	fmt.Println(s4, cap(s4)) // [2 3 5 2 3 5] 6  s4产生新的切片

	s0[0], s1[0] = 99, 789

	fmt.Println(s2[0], s3[0], s4[0]) //s2[0]=789 s3[0]=99 s4[0] 2
*/
// ----------------------------------------------------------------
/*
 var s = append([]string(nil), "array", "slice")
	fmt.Println(s)      // [array slice]
	fmt.Println(cap(s)) // 2
	s = append(s, "map")
	fmt.Println(s)      // [array slice map]
	fmt.Println(cap(s)) // 4
	s = append(s, "channel")
	fmt.Println(s)      // [array slice map channel]
	fmt.Println(cap(s)) // 4
*/

//-----------------------------------------------------
/*
var s = append([]string(nil), "格子衫来", "gzsl")
	fmt.Println(s)

	fmt.Println(cap(s)) //2

	s = append(s, "map") //没有足够的槽位、产生扩容（开辟新内存）

	fmt.Println(s)
	fmt.Println(cap(s)) //4

	s = append(s, "channel")
	fmt.Println(s)
	fmt.Println(cap(s)) //4
*/

// 4.make内置函数创建切片

/*
	  // 创建映射。
		fmt.Println(make(map[string]int)) // map[]
		m := make(map[string]int, 3)
		fmt.Println(m, len(m)) // map[] 0
		m["C"] = 1972
		m["Go"] = 2009
		fmt.Println(m, len(m)) // map[C:1972 Go:2009] 2

		// 创建切片。
		s := make([]int, 3, 5)
		fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5
		s = make([]int, 2)
		fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
*/
// ------------------------------------------------------------
/*
// 创建映射
	fmt.Println(make(map[string]int)) // map[] 非零值映射

	m := make(map[string]int, 3)
	fmt.Println(m, len(m))

	m["gzsl"] = 1996
	m["Go"] = 2009
	fmt.Println(m, len(m))

	// 创建切片
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s))
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s))
*/

// 5.使用内置new函数来创建容器值
/*
    m := *new(map[string]int)   // <=> var m map[string]int
	fmt.Println(m == nil)       // true
	s := *new([]int)            // <=> var s []int
	fmt.Println(s == nil)       // true
	a := *new([5]bool)          // <=> var a [5]bool
	fmt.Println(a == [5]bool{}) // true
*/
func main() {

	// 用new函数开辟出来的【值】均为零值
	m := *new(map[string]int)
	fmt.Println(m == nil)

	s := *new([]int)
	fmt.Println(s == nil)

	a := *new([5]bool)
	fmt.Println(a == [5]bool{})
}
