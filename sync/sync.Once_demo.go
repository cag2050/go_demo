// https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter16/16.01.md#once
package main

import "sync"

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	onceBody := func() {
		println("Only once")
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(onceBody)
		}()
	}

	wg.Wait()
}
