package main

import "fmt"

var c, python, java bool

const HOGE = "hoooooo"

func main() {
	var i int
	fmt.Println(i, c, python, java)

	// 初期化子がある場合型を省略できる
	var hoge, fuga = "hoge", 100
	fmt.Println(hoge, fuga)

	// 関数内ではvar の代わりに := を使うことができる
	aaa := 1001
	fmt.Println(aaa)

	// 定数
	fmt.Println(HOGE)

	const Truth = true
	fmt.Println(Truth)

}
