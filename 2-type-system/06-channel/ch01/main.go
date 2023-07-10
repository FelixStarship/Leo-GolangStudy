package main

import (
	"fmt"
	"time"
)

// 1.通道
/*
 package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // 一个非缓冲通道
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch    // 此操作编译不通过
		ch <- x*x  // 阻塞在此，直到发送的值被接收
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch      // 阻塞在此，直到有值发送到c
		fmt.Println(n) // 9
		// ch <- 123   // 此操作编译不通过
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	<-done // 阻塞在此，直到有值发送到done
	fmt.Println("bye")
}
*/

/*
// 通过通讯共享内存
	c := make(chan int) //非缓冲通道

	// 写请求
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		ch <- x * x
	}(c, 3)

	done := make(chan struct{})

	// 读请求
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)

	<-done // 读阻塞
	fmt.Println("格子衫来")
*/

/*
c := make(chan int, 2) // 一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok) // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok) // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c) // 此行将产生一个恐慌
	c <- 7   // 如果上一行不存在，此行也将产生一个恐慌。
*/

/*
 c := make(chan int, 2) //一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false

	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	//close(c)   //panic: close of closed channel
	//c <- 7     //panic: send on closed channel
*/

/*
    var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"   // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
*/

func main() {

	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "传球")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}

	go kickBall("坤坤")
	go kickBall("梅西")
	go kickBall("格子衫来")
	go kickBall("吴亦凡")

	ball <- "裁判"

	var c chan bool
	<-c
}
