package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go hello(i)
	}
	time.Sleep(10 * time.Second)
}

func hello(i int) {
	fmt.Printf("%d 回目\n", i)
}
