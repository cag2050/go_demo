package main

func main() {
	var b int = 15
	var a int
	arr := [6]int{11, 22, 33, 44}

	for a := 0; a < 10; a++ {
		println(a)
	}

	for a < b {
		a++
		println(a)
	}
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。
	for index, value := range arr {
		println(index, value)
	}
}
