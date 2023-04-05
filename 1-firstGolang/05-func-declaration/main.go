package main

// go函数声明

// [关键字] [函数名] [形参] [返回值]

func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s,d
}

// 1.第一部分是[func]关键字
// 2.第二部分是函数名称，函数名称必须是一个标识符
// 3.第三部分是输入参数列表
// 4.函数返回值
// 5.函数体

// 函数返回结果为【匿名】
func SquaresOfSumAndDiff1(a int64, b int64) (int64, int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
}

func SquaresOfSumAndDiff2(a int64, b int64) (s, d int64) {

	//return (a + b) * (a + b), (a - b) * (a - b)

	s = (a + b) * (a + b)
	d = (a - b) * (a - b) // 等价上面返回结果
	return s, d
}

func VersionString() string {
	return "格子衫来"
}

func CompareLower4bits(m, n uint32) (r bool) {
	r = m&0xF > n&0xf
	return
}

var v = VersionString()

func main() {

	// 函数调用
	println(v)

	x, y := SquaresOfSumAndDiff(3, 6)

	println(x, y)

	b := CompareLower4bits(uint32(x), uint32(y))

	println(b)

	// 匿名函数
	x1, y1 := func() (int, int) {
		println("this function has no 格子衫来！")
		return 3, 4
	}()

	func() {
		println(x1, y1)
	}()

}
