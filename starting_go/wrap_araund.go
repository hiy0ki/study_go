package main

import (
	"fmt"
	"math"
)

func main () {
	fmt.Printf("unit 32 max value = %d\n", math.MaxUint32)
	

	b := byte(255)
	fmt.Println(b)
	fmt.Println(b + byte(255))
	fmt.Println(b + 1)
}
