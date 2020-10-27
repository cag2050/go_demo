package main

import "fmt"

func main() {
	var key = fmt.Sprintf("%s:%d:%s:%s:%s:%d:%s:%d", "ab_refactor", 1, "req.DeviceCode", "req.MemberID", "req.Tuid", 2, "req.OS", 111)
	println(key)
}
