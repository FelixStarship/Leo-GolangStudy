package main

// 2。选择器遮挡和碰撞
/*
 type A struct {
	x string
}
func (A) y(int) bool {
	return false
}

type B struct {
	y bool
}
func (B) x(string) {}

type C struct {
	B
}
*/

type A struct {
	x string
}

func (A) y(int) bool {
	return false
}

type B struct {
	y bool
}

func (B) x(string) {}

type C struct {
	B
}

var v1 struct {
	A
	B
}

// 出现碰撞
//func f1() {
//	_ = v1.x // 不明确的引用
//	_ = v1.y
//}

var v2 struct {
	A
	C
}

// 出现遮挡 A遮挡了C
func f2() {
	_ = v2.x
	_ = v2.y
}

// 堆栈溢出

/*
type I interface {
	m()
}

type T struct {
	I
}

var t T
	var i = &t
	t.I = i
	i.m() // 将调用t.m()，然后再次调用i.m()，......
*/

type I interface {
	m()
}

type T struct {
	I
}

func main() {

	// fatal error: stack overflow
	var t T
	var i = &t
	t.I = i
	i.m() // 循环引用
}
