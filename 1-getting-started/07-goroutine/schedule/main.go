package main

import (
	"fmt"
	"runtime"
)

/*
协程调度

// linux：当前逻辑cpu的个数：cat /proc/cpuinfo| grep "processor"| wc -l
1.并非所有处于运行状态的协程都在执行。在【任一时刻】，只能最多有和【逻辑CPU数目】一样多的协程在同时执行
*/
func main() {
	// 当前系统逻辑cpu个数
	// 每个逻辑CPU在同一时刻只能最多执行一个协程
	fmt.Println(runtime.NumCPU())
	// 逻辑处理器P
	fmt.Println(runtime.GOMAXPROCS(0))
	// M-P-G模型:
	/*		M表示系统线程
		    P表示逻辑处理器（并非上述的逻辑CPU）
	        G表示协程
	*/
}
