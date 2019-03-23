package main

func main(){
	// if 支持初始化语句，可定义代码块局部变量。
	// 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
	if a := 1; a > 0 {
		println("a > 0")
	} else {
		println("a <= 0")
	}

	var b = -1
	if println("init"); b > 0 {
		println("b > 0")
	} else {
		println("b <= 0")
	}
}
