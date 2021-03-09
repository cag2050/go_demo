package main

import (
	"time"
)

func main() {
	startT := time.Now()
	timeCost := time.Since(startT)
	println("connect etcd time_cost:%v", timeCost)
}
