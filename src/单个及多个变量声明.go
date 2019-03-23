package main

import "fmt"

// 这种因式分解关键字的写法一般用于声明全局变量
// 函数外定义的变量称为全局变量，此处为全局变量
var (
	global_multi1 float64
	global_multi2 bool
)

func main() {
	// 单变量声明，有以下三种
	var first int
	var second = true
	// 这种不带声明格式的只能在函数体中出现
	third := "str"
	fmt.Println(first, second, third)

	// 多变量声明，有以下四种
	var multi11, multi12, multi13 string
	// 并行 或 同时 赋值
	var multi21, multi22, multi23 = "str21", "str22", "str23"
	// 这种不带声明格式的只能在函数体中出现
	multi31, multi32, multi33 := "str31", "str32", "str33"
	// 这种因式分解关键字的写法一般用于声明全局变量
	// 函数内定义的变量称为局部变量，此处为局部变量
	// Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑使用。
	var (
		multi41 float64
		multi42 bool
	)
	fmt.Println(multi11, multi12, multi13)
	fmt.Println(multi21, multi22, multi23)
	fmt.Println(multi31, multi32, multi33)
	fmt.Println(multi41, multi42)

	// 全局变量：允许声明但不使用
	//fmt.Println(global_multi1, global_multi2)

	// &i 来获取变量 i 的内存地址
	fmt.Println(&third)
}
