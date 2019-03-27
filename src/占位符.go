package main

import "fmt"

func main() {
	var str = "aaa"
	fmt.Printf("值为 %s，指针地址：%p", str, &str)
}