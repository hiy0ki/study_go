package main

import (
	"fmt"
)

func main() {
	fruits := [3]string{"apple", "banana", "grape"}

	for i, s := range fruits {
		// i はindex
		// s は要素
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	// rangeは複数の要素を保持する性質を備えるデータ型に使える
	// 文字列型をrangeで使うとrune型で取り出す
	strings := "あいうえお"
	for i, s := range strings {
		fmt.Printf("strings[%d]=%s %#U\n", i, s, s)
	}
}
