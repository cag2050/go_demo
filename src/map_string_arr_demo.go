package main

import "fmt"

func main() {
	Header := map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}
	println(Header)
	headerStr := ""
	for key, value := range Header {
		for _, value2 := range value {
			headerStr += key + ":" + value2 + ";"
		}
	}
	//println(headerStr)
	println(fmt.Sprintf("%v", headerStr))
}
