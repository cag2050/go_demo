package main

func main() {
	// 定义切片，2种方法：
	// 1. 可以声明一个未指定大小的数组来定义切片
	var arr1 []int
	println(len(arr1))
	// 2. 使用make()函数来创建切片
	//  len 是数组的长度并且也是切片的初始长度
	var arr2 []int = make([]int, 4)
	// 也可以简写为：
	arr3 := make([]int, 4)
	println(len(arr2))
	println(len(arr3))

	// 初始化切片
	arr4 := []int{1, 2, 3}
	arr5 := arr4[1:2]
	println(len(arr4))
	println(len(arr5))
	println(arr5[0])

	arr6 := make([]int, 3, 10)
	println(cap(arr6))
}
