package main

func main() {
	var arr1 [10] float32
	// 初始化数组中 {} 中的元素个数不能大于 [] 中的数字
	arr1 = [10]float32{1, 2, 3, 4, 5}
	println(len(arr1))

	// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
	var arr2 = []int{1, 2, 3}
	var arr3 = [...]int{1, 2, 3}
	println(len(arr2))
	println(len(arr3))


	arr4 := []string{}
	println(len(arr4))
}
