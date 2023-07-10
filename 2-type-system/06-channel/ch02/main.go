package main

// 2.遍历通道
/*
  数据接收和发送操作都属于简单语句
  package main

import (
	"fmt"
	"time"
)

func main() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // 步尾语句
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
*/

/*
	 fibonacci := func() chan uint64 {
			c := make(chan uint64)
			go func() {
				var x, y uint64 = 0, 1
				for ; y < (1 << 63); c <- y {
					x, y = y, x+y
				}
				close(c)
			}()
			return c
		}

		c := fibonacci()
		for x, ok := <-c; ok; x, ok = <-c {
			time.Sleep(time.Second)
			fmt.Println(x)
		}

		for v := range c {

		}

		for {
			v, ok := <-c
			if !ok {
				break
			}
		}
*/

/*
var c chan struct{} // nil
	select {
	case <-c:             // 阻塞操作
	case c <- struct{}{}: // 阻塞操作
	default:
		fmt.Println("Go here.")
	}
*/

/*
var c chan struct{}

	select {
	case <-c:
	case c <- struct{}{}:
	default:
		fmt.Println("格子衫来")
	}
*/

/*
	  c := make(chan string, 2)
		trySend := func(v string) {
			select {
			case c <- v:
			default: // 如果c的缓冲已满，则执行默认分支。
			}
		}
		tryReceive := func() string {
			select {
			case v := <-c: return v
			default: return "-" // 如果c的缓冲为空，则执行默认分支。
			}
		}
		trySend("Hello!") // 发送成功
		trySend("Hi!")    // 发送成功
		trySend("Bye!")   // 发送失败，但不会阻塞。
		// 下面这两行将接收成功。
		fmt.Println(tryReceive()) // Hello!
		fmt.Println(tryReceive()) // Hi!
		// 下面这行将接收失败。
		fmt.Println(tryReceive()) // -
*/

/*
c := make(chan string, 2) //
	trySend := func(v string) {
		select {
		case c <- v:
		default:

		}
	}

	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-格子衫来"
		}
	}

	trySend("格子")
	trySend("衫来")
	trySend("格子衫来")
	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
*/

/*
c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // 若此分支被选中，则产生一个恐慌
	case <-c:
	}
*/

func main() {

	// 所有case均非阻塞、【随机选择一个case执行】
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: //panic: send on closed channel
	case <-c:

	}

}
