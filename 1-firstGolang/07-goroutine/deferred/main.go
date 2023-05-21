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

func main() {

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
}
