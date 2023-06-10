package main

import "fmt"

/*
defer:延迟调用函数
    defer fmt.Println("The third line.")
	defer fmt.Println("The second line.")
	fmt.Println("The first line.")

当一个延迟调用语句被执行时，其中的延迟函数调用不会立即被执行，而是被推入由当前协程维护的一个【延迟调用队列】（一个后进先出队列）。
当一个函数调用返回（此时可能尚未完全退出）并进入它的退出阶段后，所有在执行此函数调用的过程中已经被推入延迟调用队列的调用将被按照它们被推入的顺序【逆序】被弹出队列并执行。
*/

/*
defer fmt.Println("衫来")
	defer fmt.Println("格子")
	fmt.Println("格子衫来")

	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")

	//var a int
	//defer a //defer中的表达式必须为函数  defer fn:
	return
	defer fmt.Println("not reachable")
*/

/*
// 一个延迟调用可以修改包含此延迟调用的最内层函数的返回值
package main

import "fmt"

func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值
	}()

	return n + n // <=> r = n + n; return
}

func main() {
	fmt.Println(Triple(5)) // 15
}
*/

/*
协程和延迟调用的实参的估值时刻
package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i) //0,1,2<=>2,1,0  //一个延迟调用的实参是在此调用对应的延迟调用语句被执行时被估值的。 或者说，它们是在此延迟调用被推入延迟调用队列时被估值的。 这些被估值的结果将在以后此延迟调用被执行的时候使用。
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)  // 0,1,2<=>3  //一个匿名函数体内的表达式是在此函数被【执行】的时候才会被逐渐【估值】的，不管此函数是被普通调用还是【延迟】/【协程】调用。
			}()
		}
	}()
}

*/

/*
  恐慌（panic）和恢复（recover）
  func panic(v interface{})
  func recover() interface{}

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}

*/

func main() {
	defer func() {
		fmt.Println("格子衫来，正常退出")
	}()

	fmt.Println("格子衫来")

	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()

	panic("白白!")

	// 进入函数退出阶段：恐慌产生
	fmt.Println("执行不到")

}

// fmt.Println(Triple(5))
func Triple(n int) (r int) {
	defer func() {
		r += n
	}() //<=> 10+5=15  为啥不是20?  //协程和延迟调用的实参的估值时刻
	return n + n // 10
}
