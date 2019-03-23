package main

func main()  {
	var a int = 4
	var ptr *int
	// ptr 包含了 'a' 变量的地址
	ptr = &a
	println(a)
	println(&a)
	println(*ptr)
	println(ptr)
}
