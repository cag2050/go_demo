package main

import (
	"fmt"
	"unsafe"
)

func main() {
	c := make(chan int, 3)
	fmt.Printf("%d", unsafe.Sizeof(c))
}
