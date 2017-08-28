package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println(os.Args[1])

	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	n := i + i

	fmt.Println(n)

	fmt.Println(time.Now())
}
