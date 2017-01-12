package main

import "fmt"

func main() {
	var name string = "hogehoge"
	fmt.Println(name)

	var num int = 100
	fmt.Println(num)

	var is bool = true
	fmt.Println(is)

	// 型推論
	a := "hoge"
	fmt.Println(a)
	// a := 123
	// fmt.Println(a)

	fmt.Printf("hello world\n")
}
