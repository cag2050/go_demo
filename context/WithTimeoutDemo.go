package main

import (
	"context"
	"errors"
	"time"
)

func test() (string, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	var err error
	var result string
	ch := make(chan bool, 1)
	//defer close(ch)

	// 写在这里：超时
	time.Sleep(5 * time.Second)

	go func(string, error) {
		println("goroutine 执行")
		result = "没超时"
		err = errors.New("sss")
		ch <- true
	}(result, err)

	// 写在这里：没超时
	//time.Sleep(5 * time.Second)

	// 只是比较 ctx.Done() 和 <-ch 谁先到来
	select {
	case <-ctx.Done():
		return "超时", nil
	case <-ch:
		return result, err
	}
}

func main() {
	str, err := test()
	println("str: " + str)
	if err != nil {
		println(err.Error())
	}
}
