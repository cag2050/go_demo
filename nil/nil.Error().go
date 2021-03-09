package main

func main() {
	var err error = nil
	/*
		下面这行代码报 panic：
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1056f3f]
	*/
	println(err.Error())
}
