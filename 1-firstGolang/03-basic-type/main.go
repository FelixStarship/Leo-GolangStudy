package main

/*基本类型和它们的字面量表示*/
/*
Go支持如下内置基本类型：
一种内置布尔类型：bool。
11种内置整数类型：int8、uint8、int16、uint16、int32、uint32、int64、uint64、int、uint和uintptr。
两种内置浮点数类型：float32和float64。
两种内置复数类型：complex64和complex128。
一种内置字符串类型：string。
*/

/*type关键字自定义类型*/

func main() {
	//package builtin  内置类型

	/*

		Go中有两种内置类型别名（type alias）：
		byte是uint8的内置别名。 我们可以将byte和uint8看作是同一个类型。
		rune是int32的内置别名。 我们可以将rune和int32看作是同一个类型。
		type byte = uint8
		type rune = int32

	*/

	/*自定义类型别名*/
	type boolean = bool

	type Text = string // TEXT和string是不同类型

}
