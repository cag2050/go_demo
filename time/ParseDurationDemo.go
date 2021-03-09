package main

import "time"

func main() {
	var timeoutStr = "100ns"
	// int64的Nanosecond纳秒
	var timeout time.Duration
	timeout, err := time.ParseDuration(timeoutStr)
	// 转换错误
	if err != nil {
		println(err.Error())
		// 设置默认值
		timeout = 300 * time.Nanosecond
	}
	println(timeout)
}
