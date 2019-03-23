package main

import (
	"fmt"
	"time"
)

func main() {
	tickTimer := time.NewTicker(1 * time.Second)
	barTimer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-tickTimer.C:
			fmt.Println("tick")
		case <-barTimer.C:
			fmt.Println("bar")
		}
	}
}
