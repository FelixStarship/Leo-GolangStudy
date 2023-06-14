package main

import "fmt"

// 1.GO容器类型
/*
// 下面这些切片字面量都是等价的。
[]string{"break", "continue", "fallthrough"}
[]string{0: "break", 1: "continue", 2: "fallthrough"}
[]string{2: "fallthrough", 1: "continue", 0: "break"}
[]string{2: "fallthrough", 0: "break", "continue"}

// 下面这些数组字面量都是等价的。
[4]bool{false, true, true, false}
[4]bool{0: false, 1: true, 2: true, 3: false}
[4]bool{1: true, true}
[4]bool{2: true, 1: true}
[...]bool{false, true, true, false}
[...]bool{3: false, 1: true, true}


// 1.容器类型初始化
// 1.初始化数组
	a := [4]bool{false, true, true, false}
	// 2.初始化切片*****
	a1 := []bool{0: false, 1: true, 2: false, 3: false}
	// 3.初始化数组,编译器推断长度
	a2 := [...]string{"格子衫来", "格子衫来", "格子衫来"}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
*/

// 2.容器类型零值的字面量表示形式
/*

 1.*****指针一样，所有【切片】和【映射】类型的【零值】均用预声明的标识符【nil】来表示。******

 2.容器字面量是不可寻址的但可以被取地址
 package main

import "fmt"

func main() {
	pm := &map[string]int{"C": 1972, "Go": 2009}
	ps := &[]string{"break", "continue"}
	pa := &[...]bool{false, true, true, false}
	fmt.Printf("%T\n", pm) // *map[string]int
	fmt.Printf("%T\n", ps) // *[]string
	fmt.Printf("%T\n", pa) // *[4]bool
}

---------------------------------
// 容器字面量取地址
	pm := &map[string]int{"格子衫来": 1996, "GO": 2009}
	ps := &[]string{"格子衫来", "格子衫来"}
	pa := &[...]bool{false, true, false}

	fmt.Printf("%T\n", pm) //*map[string]int
	fmt.Printf("%T\n", ps) //*[]string
	fmt.Printf("%T\n", pa) //*[3]bool
*/

// 3.内嵌组合字面量可以被简化
/*
  // heads为一个切片值。它的类型的元素类型为*[4]byte。
// 此元素类型为一个基类型为[4]byte的指针类型。
// 此指针基类型为一个元素类型为byte的数组类型。
var heads = []*[4]byte{
	&[4]byte{'P', 'N', 'G', ' '},
	&[4]byte{'G', 'I', 'F', ' '},
	&[4]byte{'J', 'P', 'E', 'G'},
}
-------------------------
type language struct {
	name string
	year int
}

var _ = [...]language{
	language{"C", 1972},
	language{"Python", 1991},
	language{"Go", 2009},
}
---------------------------------
type LangCategory struct {
	dynamic bool
	strong  bool
}


var _ = map[LangCategory]map[string]int{
	LangCategory{true, true}: map[string]int{
		"Python": 1991,
		"Erlang": 1986,
	},
	LangCategory{true, false}: map[string]int{
		"JavaScript": 1995,
	},
	LangCategory{false, true}: map[string]int{
		"Go":   2009,
		"Rust": 2010,
	},
	LangCategory{false, false}: map[string]int{
		"C": 1972,
	},
}

*/
func main() {

	type LangCategory struct {
		dynamic bool
		strong  bool
	}

	// key:value
	var lang = map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"格子衫来": 1996,
		},
		LangCategory{true, true}: map[string]int{
			"格子衫来": 1996,
		},
		LangCategory{true, true}: map[string]int{
			"格子衫来": 1996,
		},
	}

	fmt.Println(lang)

	var lang1 = map[LangCategory]map[string]int{
		{true, true}:  {"格子衫来": 1996},
		{false, true}: {"格子衫来": 1996},
	}

	fmt.Println(lang1)
	// 简化之前
	var heads = []*[4]byte{
		&[4]byte{'G', 'Z', 'S', 'L'},
		&[4]byte{'G', 'Z', 'S', 'L'},
		&[4]byte{'G', 'Z', 'S', 'L'},
	}

	// 简化之后
	var heads1 = []*[4]byte{
		{'G', 'Z', 'S', 'L'},
		{'G', 'Z', 'S', 'L'},
		{'G', 'Z', 'S', 'L'},
	}

	fmt.Println(heads)
	fmt.Println(heads1)

	type language struct {
		name string
		year int
	}

	var _ = [...]language{
		language{"格子衫来", 1996},
		language{"GO", 1996},
		language{"Python", 1996},
	}

	var _ = [...]language{
		{"格子衫来", 1996},
	}
}
