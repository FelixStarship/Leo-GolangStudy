package main

import (
	"sync"
	"time"
)

/*

1.协程的状态：

注意事项：
1.一个活动中的协程可以处于两个状态：运行状态和阻塞状态
2.当一个新协程被创建的时候，它将自动进入运行状态，一个协程只能从运行状态而不能从阻塞状态退出。
如果因为某种原因而导致某个协程一直处于阻塞状态，则此协程将永远不会退出。
3.一个处于阻塞状态的协程不会自发结束阻塞状态，它必须被另外一个协程通过某种并发同步方法来被动地结束阻塞状态。
4. 如果一个运行中的程序当前【所有的协程】都出于【阻塞状态】，则这些协程将永远阻塞下去，程序将被视为死锁了。
当一个程序死锁后，官方标准编译器的处理是让这个程序崩溃


*/
/*
package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait() // 阻塞在此
	}()
	wg.Wait() // 阻塞在此
}
*/

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait() // 子协程进去阻塞
	}()

	wg.Wait()
}
