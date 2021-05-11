package main

import (
	"fmt"
	"time"
)

func main() {
	requests := []int{12, 2, 3, 41, 5, 6, 1, 12, 3, 4, 2, 31}
	for _, request := range requests {
		go run(request) //开启多个协程
	}

	ch := make(chan struct{})
	go func(a chan struct{}) {
		time.Sleep(5 * time.Second)
		a <- struct{}{}
	}(ch)
	select {
	case <-ch:

	}
}

func run(num int) {
	defer func() {
		if err := recover();err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	if num%5 == 0 {
		panic("请求出错la")
	}
	fmt.Printf("%d beng sha ka la ka !\n", num)
}