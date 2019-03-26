package main

import (
	"fmt"
)

func main() {
	var a interface{} = "55"
	// type 只能用在 switch 语句中
	switch vv := a.(type) {
	// interface{} 匹配所有类型；switch 找到第一个成功的 case，执行完就会退出。
	case interface{}:
		fmt.Println("interface{}")
	case string:
		fmt.Printf("vv type = %T, vv value = %+v, a type = %T, a value = %+v\n", vv, vv, a, a)
	}

	var marks = 90
	switch marks {
	case 90:
		println("A")
	case 80:
		println("B")
	case 50, 60, 70:
		println("C")
	default:
		println("D")
	}
}
