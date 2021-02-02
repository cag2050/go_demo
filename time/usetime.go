package main

import (
	"fmt"
	"time"
)

func sum(n int) int {
	startT := time.Now()		//计算当前时间

	total := 0
	for i:=1; i <= n; i++ {
		total += i
	}

	timeCost := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", timeCost)
	return total
}

func main() {
	count := sum(100)
	fmt.Printf("count = %v\n", count)
}