package main

import (
	"fmt"
	"time"
)

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done")
}

func main() {
	ch := make(chan int, 20)

	go receive("1st. goruchine", ch)
	go receive("2st. goruchine", ch)
	go receive("3st. goruchine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

	close(ch)

	time.Sleep(3 * time.Second)
}
