package main

import "fmt"

func main() {
	key := fmt.Sprintf("%d%s%s", 1, ":", "1234")
	println(key)
}
