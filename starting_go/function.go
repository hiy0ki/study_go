package main

import (
	"fmt"
)

func main() {
	fmt.Println(plus(10, 45))

	hello()

	a, b := div(20, 32)
	fmt.Println(a, b)

	// 戻り値の破棄
	a2, _ := div(20, 32)
	fmt.Println(a2)

	fmt.Println(hoge())
	fmt.Println(hoge2())

	fmt.Println(div2(4, 3))

	// 無名関数 (クロージャ)
	f := func(x, y int) int { return x + y }
	fmt.Println(f(4, 3))
	// 無名関数の型定義の確認
	fmt.Printf("%T\n", f)

	// 関数の代入
	var f_alias = f // f_alias := f でも同じ
	fmt.Printf("%T\n", f_alias)

	// 関数を返す関数
	returnFunc()()

	// 関数を引数に取る関数
	callFunc(returnFunc())
}

// ベーシックな関数定義
func plus(x, y int) int {
	return x + y
}

// 戻りの型定義を省略したらvoidになる
func hello() {
	fmt.Println("hello world")
	return // 戻り値がない場合はあってもなくても
}

// 複数の戻り値
func div(x, y int) (int, int) { // 戻り値の型定義のカッコは省略できない
	a := x + y
	b := x - y

	return a, b
}

/*
  戻り値を表す変数
*/
// 関数内のローカル変数定義の省略
func hoge() (a int) {
	return
}

// hogeと同じ
func hoge2() int {
	var a int
	return a
}

// div()のローカル変数定義を省略してみる
func div2(x, y int) (a, b int) {
	a = x + y
	b = x - y

	return a, b
}

// 関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// 関数を引数に取る関数
func callFunc(f func()) {
	f()
}
