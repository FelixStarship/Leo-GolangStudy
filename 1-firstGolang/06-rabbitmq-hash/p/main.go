package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 1)
	timer2 := time.NewTimer(time.Second * 5)

exit:
	for {
		select {
		case <-timer1.C:
			fmt.Println(hash("my", 3))
			timer1.Reset(time.Second * 1)
		case <-timer2.C:
			break exit
		}
	}
}

// 哈希函数，将字符串哈希为一个介于 0 和 m-1 之间的整数
func hash(s string, m int) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h*31 + int(s[i])) % m
	}
	return h
}
