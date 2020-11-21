package main

import "fmt"

func main() {
	// 同步通道a，异步通道b
	a, b := make(chan int), make(chan int, 3)

	b <- 1
	b <- 2

	fmt.Printf("len(a): %d, cap(a): %d\n", len(a), cap(a)) // len和cap，对于同步通道都返回0，据此可判断通道是同步。
	fmt.Printf("len(b): %d, cap(b): %d\n", len(b), cap(b))
}
