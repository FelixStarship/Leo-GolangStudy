package main

import "fmt"

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

/*
  words := []string{
		"Go", "is", "a", "high",
		"efficient", "language.",
	}

	// fmt.Println函数的原型为：
	// func Println(a ...interface{}) (n int, err error)
	// 所以words...不能传递给此函数的调用。

	// fmt.Println(words...) // 编译不通过

	// 将[]string值转换为类型[]interface{}。
	iw := make([]interface{}, 0, len(words))
	for _, w := range words {
		iw = append(iw, w)
	}
	fmt.Println(iw...) // 编译没问题
*/

/*
 words := []string{
		"gzsl", "Go", "a", "格子衫来",
		"长安九万里", "封神榜",
	}

	//fmt.Println(words....)  // []interface{}

	iw := make([]interface{}, 0, len(words))

	// 追加元素
	for _, w := range words {
		iw = append(iw, w)
	}

	fmt.Println(iw...) // 编译ok
*/

/*
 type I interface {
	m(int)bool
}

type T string
func (t T) m(n int) bool {
	return len(t) > n
}
var i I = T("gopher")
	fmt.Println(i.m(5))                          // true
	fmt.Println(I.m(i, 5))                       // true
	fmt.Println(interface {m(int) bool}.m(i, 5)) // true

	// 下面这几行被执行的时候都将会产生一个恐慌。
	I(nil).m(5)
	I.m(nil, 5)
	interface {m(int) bool}.m(nil, 5)
*/

type I interface {
	M(int) bool
}

type T string

func (t T) M(n int) bool {
	return len(t) > n
}

func main() {

	var i I = T("gzsl") // 将一个字符串"gzsl"转换为T类型

	fmt.Println(i.M(5))

	fmt.Println(I.M(i, 5))

	fmt.Println(interface{ M(int) bool }.M(i, 5))

	//I(nil).M(5) //panic: runtime error: invalid memory address or nil pointer dereference
	///I.M(nil, 5)
	//interface{ M(int) bool }.M(nil, 5)
}
