package main

import "fmt"

func main() {
	// 带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

	// 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
	// 缓冲区大小为2
	channel := make(chan int, 2)

	// 因为 channel 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	channel <- 1
	channel <- 2

	// 获取这两个数据
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
