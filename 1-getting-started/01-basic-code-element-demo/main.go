package main //指定当前源文件所在的包名
import "math/rand"

const MaxRand = 16 //声明一个具名整型常量

func StatRandomNumbers(numRands int) (int, int) {
	// 声明了两个变量（类型都为int，初始值都为0）
	var a, b int
	for i := 0; i < numRands; i++ {
		// if-else条件代码控制块
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++
		}
	}
	return a, b
}

func main() {
	var num = 100
	x, y := StatRandomNumbers(num)
	print("Result:", x, "+", y, "=", num, "? ")
	println(x+y == num)
}
