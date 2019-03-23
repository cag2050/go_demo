package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main()  {
	// Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
	// goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
	// 同一个程序中的所有 goroutine 共享同一个地址空间。
	// 执行代码，你会看到输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 线程在执行
	go say("world")
	say("hello")
}
