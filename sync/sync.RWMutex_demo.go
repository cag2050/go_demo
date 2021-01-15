// https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter16/16.01.md

package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁分为读锁和写锁，读数据的时候上读锁，写数据的时候上写锁。
// 有写锁的时候，数据不可读不可写。有读锁的时候，数据可读，不可写。
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