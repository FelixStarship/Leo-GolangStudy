package main

import "fmt"

// 第六讲：golang中的基本流程控制语句
/*
  if else 条件分支代码块
  for 循环代码块
  switch-case多条件分支代码块
*/

/*
    // break 语句,用来跳出包含此break语句的最内层for循环
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

    // continue 语句,一条continue语句可以被用来提前结束包含此continue语句的最内层for循环的当前循环步
    for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

    // switch-case
    switch InitSimpleStatement; CompareOperand0 {
	case CompareOperandList1:
		// do something
	case CompareOperandList2:
		// do something
	...
	case CompareOperandListN:
		// do something
	default:
		// do something
	}

    rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println(n, "is a multiple of 9.")
	case 1, 2, 3:
		fmt.Println(n, "mod 9 is 1, 2 or 3.")
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}

    // fallthrough
    rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n=n", n)
		fallthrough
	case 5, 6, 7, 8:
		n := 99
		fmt.Println("n=n", n)
		fallthrough
	default:
		fmt.Println("n=n", n)
	}

    goto跳转标签和跳转语句
    i := 0

Next: // 跳转标签声明 【一条跳转标签声明之后必须立即跟随一条语句。 如果此声明的跳转标签使用在一条goto语句中，则当此条goto语句被执行的时候，执行将跳转到此跳转标签声明后跟随的语句。】
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next // 跳转
	}

*/

func main() {

	fmt.Println(FindSmallestPrimeLargerThan(1))
}

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}
