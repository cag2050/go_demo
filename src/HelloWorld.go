// 必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

import "fmt"

// main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
func main() {
	// 单行注释
	/*
		多行注释（也叫：块注释）
		多行注释
	*/
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")
}
