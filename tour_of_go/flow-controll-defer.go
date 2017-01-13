package main

import "fmt"

func main(){
	// 呼び出し元の関数がreturnするタイミングで実行される
	defer fmt.Println("world")

	fmt.Println("hello")

	// deferへ渡す関数が複数の場合LIFOで実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
