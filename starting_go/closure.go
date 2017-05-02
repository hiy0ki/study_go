package main

import (
	"fmt"
)

func later() func(string) string {
	// 一つ前に与えられた文字列を保存するための変数
	var store string

	// 引数に文字列を取り保存されていた文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func increment() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main () {
	f := later()

	fmt.Println(f("hoge"))
	fmt.Println(f("is"))
	fmt.Println(f("great!"))

	f2 := increment()

	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())

	f3 := increment()
	fmt.Println(f3()) // クロージャ間で共有されるわけではないので注意

}
