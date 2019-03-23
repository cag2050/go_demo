package main

func main() {
	// 遍历数组
	arr := []int{11, 22, 33}
	for index, value := range arr {
		println(index, value)
	}

	// 遍历map
	map1 := map[string]string{"a": "apple","b":"banana"}
	for key, value := range map1 {
		println(key, ":", value)
	}

	// 遍历字符串，返回字符的 Unicode值
	str := "abc"
	for index, value := range str {
		println(index, value)
	}
}
