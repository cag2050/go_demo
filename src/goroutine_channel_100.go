package main

func produce(c chan int, num int) {
	for i := 0; i < num; i++ {
		c <- i
	}
	close(c)
}

func main() {
	var chan1 = make(chan int, 100)
	go produce(chan1, 100)
	for v := range chan1 {
		println(v)
	}

}
