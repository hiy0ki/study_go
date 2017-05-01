package main

import (
	"fmt"
	"math"
)

func main () {
	// 最大値の定数
	fmt.Printf("unit 32 max value = %d\n", math.MaxUint32)
	

	b := byte(255)
	fmt.Println(b)
	fmt.Println(b + 1) // == 0
	fmt.Println(b + byte(255)) // -1したのと同じ
}
