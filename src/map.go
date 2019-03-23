package main

func main() {
	//  Map 是使用 hash 表来实现的。
	// 定义 Map，2种方式：
	// 1. 使用 map 关键字
	map1 := map[string]string{
		"a": "aaa",
		"b": "bbb",
	}

	for key, value := range map1 {
		println(key, ":", value)
	}

	// 2. 使用内建函数 make
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["French"] = "Paris"
	delete(countryCapitalMap, "French")
	for key, value := range countryCapitalMap {
		println(key, ":", value)
	}
}
