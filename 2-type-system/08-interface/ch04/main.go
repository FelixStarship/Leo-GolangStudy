package main

// 5.类型断言

/*
	 // 编译器将把123的类型推断为内置类型int。
		var x interface{} = 123

		// 情形一：
		n, ok := x.(int)
		fmt.Println(n, ok) // 123 true
		n = x.(int)
		fmt.Println(n) // 123

		// 情形二：
		a, ok := x.(float64)
		fmt.Println(a, ok) // 0 false

		// 情形三：
		a = x.(float64) // 将产生一个恐慌
*/

/*
var x interface{} = 123

	// 场景一：
	n, ok := x.(int)
	fmt.Println(n, ok)
	n = x.(int)
	fmt.Println(n)

	// 场景二：
	a, ok := x.(float64)
	fmt.Println(a, ok)

	// 场景三：    // panic: interface conversion: interface {} is int, not float64
	a = x.(float64)
*/

// 接口类型断言

/*type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

var x interface{} = DummyWriter{}
var y interface{} = "格子衫来"
var w Writer
var ok bool

w, ok = x.(Writer)
fmt.Println(w, ok) // {} true

w, ok = y.(Writer)
fmt.Println(w, ok) //nil  false
w = y.(Writer)     // panic: interface conversion: string is not main.Writer: missing method Write
*/

/*
 values := []interface{}{
		456, "gzsl", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}

	for _, x := range values {
		switch v := x.(type) {
		case []int:
			fmt.Println("int Slice:", v)
		case string:
			fmt.Println("string:", v)
		case int, float64, int32:
			fmt.Println("number:", v)
		case nil:
			fmt.Println(v)
		default:
			fmt.Println("others:", v)
		}
	}
*/

/*
 values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}

	for _, x := range values {
		if v, ok := x.([]int); ok {
			fmt.Println("int slice:", v)
		} else if v, ok := x.(string); ok {
			fmt.Println("string:", v)
		} else if x == nil {
			v := x
			fmt.Println(v)
		} else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32 {
				v := x
				fmt.Println("number:", v)
			} else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}
*/

func main() {

}
