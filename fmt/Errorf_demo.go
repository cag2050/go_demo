package main

import "fmt"

func main() {
	err := fmt.Errorf("err str")
	if err != nil {
		//println(err.Error())
	}
}
