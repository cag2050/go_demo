package main

func main() {

	var i *int // i 是指针
	var j int
	i = &j // 取 j 的内存地址，赋值给指针
	println(i)


	type AA struct {
		a int
	}
	var a *AA
	var b AA
	a = &b
	println(a)
}
