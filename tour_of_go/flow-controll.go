package main

import (
	"fmt"
	"math"
)

func main() {

	// 繰り返し制御はfor文しかない
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 初期化と後処理は 省略できる
	sum2 := 1
	for sum2 < 10000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// セミコロンも省略できる whileと同じ
	sum3 := 1
	for sum3 < 100 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	// 無限ループはこれだけ
	// for {
	// {

	// if文もカッコは不要 中括弧は必要
	x := 3
	if x < 5 {
		fmt.Println("5 or less")
	}

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

// if文は条件の前にstatementをつけることができる
// スコープはif分内のみになる
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// この ｖ は else文のなかでも使える
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
