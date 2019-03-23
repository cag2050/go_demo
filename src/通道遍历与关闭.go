package main

import "fmt"

func fibonacci(n int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		channel <- x
		// 不明白下面这行语句
		x, y = y, x+y
	}
	close(channel)
}

func main() {
	channel := make(chan int, 10)
	go fibonacci(cap(channel), channel)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range channel {
		fmt.Println(i)
	}
}
