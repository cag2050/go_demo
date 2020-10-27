package main

import "strings"

func main()  {
	var strSlice []string

	strSlice = append(strSlice, "12")
	strSlice = append(strSlice, "34")

	str := strings.Join(strSlice, ",")
	println(str)
}
