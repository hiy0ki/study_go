package main

import (
	"fmt"
)

func main () {
	// 要素のサイズが５であるint型の配列の定義
	a := [5]int{ 1,3,4,5,6 }

	fmt.Printf("%v\n", a )

	// 要素の指定は一般的。0始まり
	fmt.Println(a[1]) // == 3

	// 以下bとcは同じ 初期値の0で埋められる
	var b [5]int
	c := [5]int{}

	fmt.Println(b)
	fmt.Println(c)

	// 初期値について、bool型はfalse string型は""(空文字)になる
	d := [3]bool{}
	e := [3]string{}

	fmt.Println(d)
	fmt.Println(e)

	// 使いみちはわからないが要素数0の配列型の定義もできる
	fmt.Println([0]int{})


	// 要素数の省略
	a1 := [...]int{1,2,3} // == [3]int{1,2,3}
	fmt.Println(a1)

	// 要素への代入
	a1[2] = 10
	fmt.Println(a1)

	/*
	配列型の代入

	要素数と要素の数が同一の変数であれば相互に代入が可能
	注意点は、参照ではなくコピーになること
	*/
	ar1 := [3]int{1,2,3}
	ar2 := [3]int{4,5,6}
	
	ar1 = ar2

	fmt.Println(ar1) // == [4,5,6]

	ar2[0] = 100

	fmt.Println(ar1) // == [4,5,6]
	fmt.Println(ar2) // == [10,5,6]

	// これはサイズ、型が違うのでコンパイルエラー
	//ar1 = [3]string{}

}
