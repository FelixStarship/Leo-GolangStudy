package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var x struct {
		a int64
		b bool
		c string
	}

	fmt.Println(unsafe.Sizeof(x.c))

}
