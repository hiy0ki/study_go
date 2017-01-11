package main

import "fmt"


func main() {
    fmt.Println(add(23,49))
    fmt.Println(swap("hello","world"))

    a, b := swap("hello","world")
    fmt.Println(a, b)
}


func add(x, y int) int {
// add(x int, y int) を省略下のが上
    return x + y
}


// 関数は複数の値を返却することができる
func swap(x, y string) (string, string) {
    return y, x
}


// naked return ステートメント
// 戻り値となる変数に名前をつける
// TODO 実装例 https://go-tour-jp.appspot.com/basics/7
