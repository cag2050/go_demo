package main

import "fmt"

func main() {
	str := "aaa"
	fmt.Printf("值为： %s，指针地址：%p, 值类型为：%T \n", str, &str, str)

	num := 10
	fmt.Printf("值为：%d，值类型为：%T \n", num, num)
}