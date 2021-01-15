package main

import "sync"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		println(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		println(2)
	}()

	wg.Wait()
}
