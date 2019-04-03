package main

import "fmt"

func myitoa (i int) string {
	str := fmt.Sprintf("%d", i)
	return str
}

func main() {
	print(myitoa(12222))
}
