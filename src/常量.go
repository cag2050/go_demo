package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	const str = "const value"
	fmt.Println(str)

	// 常量还可以用作枚举：
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown, Female, Male)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
	const (
		a = "abc"
		b = len(a)
		// 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a, b, c)
/*
	iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	const (
		// 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
		a2 = iota
		b2
		c2
	)
	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)
}
