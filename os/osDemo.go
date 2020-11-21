package main

import "os"

func main() {
	println("会执行")
	// runtime.Goexit() 立即终止当前任务，运行时会确保所有已注册延迟调用被执行。
	// os.Exit(code int) 可终止进程，但不会执行延迟调用。
	os.Exit(1)
	println("不会执行")




}
