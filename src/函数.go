package main

func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

func main()  {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	println(ret)

	c, d := swap("john", "jeff")
	println(c, d)
}