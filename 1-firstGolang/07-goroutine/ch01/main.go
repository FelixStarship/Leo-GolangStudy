package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1. go中的协程（并发和并行）

/*


  // 当一个程序的主协程退出后，此程序也就退出了，即使还有一些其它协程在运行。
  package main

import (
	"log"
	"math/rand"
	"time"
)

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	time.Sleep(2 * time.Second)
}

*/

/*
  // 并发同步
package main

import (
	"log"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成。
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // 注册两个新任务。
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成。
}
*/
//运行这个修改后的程序，我们将会发现所有的20条问候语都将在程序退出之前打印出来。

var wg sync.WaitGroup

func main() {

	rand.Seed(time.Now().UnixNano())

	log.SetFlags(0)

	wg.Add(2)
	go SayGreetings("格子", 10)
	go SayGreetings("衫来", 10)

	wg.Wait() //阻塞调用
}

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 0-2.5秒
	}
	wg.Done() //-1
}
