package main

import "errors"

func main() {
	err := errors.New("err string")
	if err != nil {
		println(err.Error())
	}
}
