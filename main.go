package main

import (
	"fmt"
	"time"
)

func main() {
	stopc := make(chan struct{})

	select {
	case <- stopc:
		fmt.Println("hello")
	case <- time.After(5 * time.Second):
		fmt.Println("after 哈哈哈")
	}
}