package main

import "fmt"

func main() {
	list := [...]int{1, 2, 3}
	println(fmt.Sprintf("%+v", "featureData"))
	println(fmt.Sprintf("%+v", list))
}
