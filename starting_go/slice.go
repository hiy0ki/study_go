package main

import (
	"fmt"
)
/*
参照型 は 3種類
1. slice
2. map
3. channel

ここではslice
*/
func main() {

	// int型のslice
	// var s []int // 宣言
	s := make([]int, 10)

	// 配列と似ているが、型として別物
	fmt.Println(s)
	fmt.Println(len(s))

	s2 := make([]int, 5, 10)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// makeを使用しないスライスの生成方法
	s3 := []int{1,2,3,4,5}
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	// 簡易スライス式
	a := [5]int{1,2,3,4,5}
	s4 := a[0:2]
	fmt.Println(s4)

	s4 = append(s4, 5)
	fmt.Println(s4)
	fmt.Println(a)

	// スライスと可変長引数
	fmt.Println(sum(1,2,3,4))
	fmt.Println(sum(1,2,3,4,6,7,8,98))
}


// 可変長引数サンプル
// 可変長引数は引数の末尾に一つだけ定義できる
func sum(s ...int) int {
	n := 0
	for _,v := range s {
		n += v
	}
	return n
}

