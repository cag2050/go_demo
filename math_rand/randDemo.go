package main

import (
	"math/rand"
	"strconv"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	var str  = "col_name" + "_random" + strconv.Itoa(rand.Intn(100) + 1)
	println(str)
}
