package main

import (
	"fmt"
)

func main() {
	/*
	 intarface{}型は、
	 - すべての型と互換性を持つ(javaのobject型的な?)
	 - 初期値としてnil(具体的な値を持っていない状態)をもつ(javaのnull的な?)
	*/
	var x interface{}
	fmt.Printf("%v", x)

	// 色々代入できる
	x = 123
	fmt.Printf("%v", x)

	x = "hogehoge"
	fmt.Printf("%v", x)

	x = 3.14
	fmt.Printf("%v", x)

	// ただし演算はできない
	var a, b interface{}
	a, b = 1, 2
	fmt.Println(a + b) // エラー
}
