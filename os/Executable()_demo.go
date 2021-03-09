package main

import "os"

func main() {
	exec, err := os.Executable()
	println(exec)
	if err != nil {
		println(err.Error())
	}
}
