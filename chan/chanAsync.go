package main

func main() {
	// 同步模式必须有配对操作的 goroutine 出现，否则会一直阻塞；
	// 异步模式在缓冲区未满或数据未读完前，不会阻塞。
	// 下面是异步模式的例子
	c := make(chan int, 3)

	c <- 1
	c <- 2

	println(<-c)
	println(<-c)
}
