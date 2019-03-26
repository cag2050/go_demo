package main

func main()  {
	// defer 的函数，参数会立即求值；
	defer println("defer ")
	println("hello")
}
