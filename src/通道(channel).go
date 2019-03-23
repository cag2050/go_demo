package main

import "fmt"

func sum(arr []int, channel chan int) {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	channel <- sum // 把 sum 发送到通道
}

func main() {
	// 通道（channel）是用来传递数据的一个数据结构。
	// 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
	// 默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端接收数据。
	arr := []int{7, 2, 8, -9, 4, 0}

	// 声明一个通道，使用chan关键字，通道在使用前必须先创建
	channel1 := make(chan int)
	go sum(arr[len(arr)/2:], channel1)
	go sum(arr[:len(arr)/2], channel1)
	x, y := <-channel1, <-channel1 // 从通道中接收

	fmt.Println(x, y, x+y)
}
