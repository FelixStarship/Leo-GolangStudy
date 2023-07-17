package main

import "fmt"

// 2.值包裹

/*
	 var i interface{}
		i = []int{1, 2, 3}
		fmt.Println(i) // [1 2 3]
		i = map[string]int{"Go": 2012}
		fmt.Println(i) // map[Go:2012]
		i = true
		fmt.Println(i) // true
		i = 1
		fmt.Println(i) // 1
		i = "abc"
		fmt.Println(i) // abc

		// 将接口值i中包裹的值清除掉。
		i = nil
		fmt.Println(i) // <nil>
*/
func main() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i)

	i = map[string]int{"格子衫来": 20230716}
	fmt.Println(i)

	i = true
	fmt.Println(i)

	i = 1
	fmt.Println(i)

	i = "gzsl"
	fmt.Println(i)

	i = nil
	fmt.Println(i)

}
