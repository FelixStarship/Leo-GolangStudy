package main

import "fmt"

func main() {

	fmt.Println(hash("my", 3))

}

// 哈希函数，将字符串哈希为一个介于 0 和 m-1 之间的整数
func hash(s string, m int) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h*31 + int(s[i])) % m
	}
	return h
}
