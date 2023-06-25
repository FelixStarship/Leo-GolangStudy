package main

import "fmt"

/*
type Person struct {
		name string
		age  int
	}
	persons := [2]Person {{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 此修改将不会体现在这个遍历过程中，
		// 因为被遍历的数组是persons的一个副本。
		persons[1].name = "Jack"

		// 此修改不会反映到persons数组中，因为p
		// 是persons数组的副本中的一个元素的副本。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
*/
// -----------------------------------
/*
type Person struct {
		name string
		age  int
	}

	persons := []Person{{"gzsl", 29}, {"格子衫来", 29}}

	// 副本
	for i, p := range persons {
		fmt.Println(i, p)
		// 切片和副本共享容器元素
		persons[1].name = "Go语言入门到进阶"
		// 循环变量的修改不会影响容器元素本身
		p.age = 31
	}

	fmt.Println("persons:", persons)
*/

/*
	  langs := map[struct{ dynamic, strong bool }]map[string]int{
			{true, false}:  {"JavaScript": 1995},
			{false, true}:  {"Go": 2009},
			{false, false}: {"C": 1972},
		}
		// 此映射的键值和元素类型均为指针类型。
		// 这有些不寻常，只是为了讲解目的。
		m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
		for category, langInfo := range langs {
			m0[&category] = &langInfo
			// 下面这行修改对映射langs没有任何影响。
			category.dynamic, category.strong = true, true
		}
		for category, langInfo := range langs {
			fmt.Println(category, langInfo)
		}

		m1 := map[struct{ dynamic, strong bool }]map[string]int{}
		for category, langInfo := range m0 {
			m1[*category] = *langInfo
		}
		// 映射m0和m1中均只有一个条目。
		fmt.Println(len(m0), len(m1)) // 1 1
		fmt.Println(m1) // map[{true true}:map[C:1972]]
*/

// ------------------------------------------------------------

/*
 langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, true}:   {"格子衫来": 1996},
		{true, false}:  {"gzsl": 1997},
		{false, false}: {"Go": 1998},
	}

	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}

	for category, langInfo := range langs {
		// 循环变量（临时变量）每次迭代都会被初始化
		m0[&category] = &langInfo // m0 映射产生了2次修改
		// 修改循环变量对lang不会产生任何影响
		category.dynamic, category.strong = true, false
	}

	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}

	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}

	fmt.Println(len(m0), len(m1))
	fmt.Println(m1)
*/

// 把数组指针当做数组来使用
/*
 var a [100]int

	for i, n := range &a { // 复制一个指针的开销很小
		fmt.Println(i, n)
	}

	for i, n := range a[:] { // 复制一个切片的开销很小
		fmt.Println(i, n)
	}
*/

/*
 var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	for i, n := range p { // panic
		fmt.Println(i, n)
	}
*/

func main() {

	var p *[5]int // p==nil
	var a [100]int
	fmt.Println(p)
	fmt.Println(a) //[0 0 0 ....]
	for i, _ := range p {
		fmt.Println(i)
	}

	for i, n := range p {
		fmt.Println(i, n)
	}

	for i, n := range &a { // 数组指针
		fmt.Println(i, n)
	}

	for i, n := range a[:] { // 子切片
		fmt.Println(i, n)
	}

}
