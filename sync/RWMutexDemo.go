package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex
var val = 0

func main() {
	m = new(sync.RWMutex)
	go read(1)
	go write(2)
	go read(3)
	time.Sleep(5 * time.Second)
}

func read(i int) {
	m.RLock()
	time.Sleep(1 * time.Second)
	fmt.Println("val:", val)
	time.Sleep(1 * time.Second)
	m.RUnlock()
}

func write(i int) {
	m.Lock()
	val = 10
	time.Sleep(1 * time.Second)
	m.Unlock()
}