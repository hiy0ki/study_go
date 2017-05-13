package main

import (
	"fmt"
)

func main() {
	var p *int // int型のポインタの宣言

	fmt.Println(p) // == nil

	v := 2
	fmt.Println(&v) // vのアドレス

	p = &v
	fmt.Println(p) // vのアドレスと同じ

	// *をつけるとデリファレンス(ポインタ型を通してデータ本体を参照すること)できる
	fmt.Println(*p) // == 2

	// pを通してvの値を更新
	*p = 5
	fmt.Println(v) // == 5

	// 関数を使ったポインタの挙動確認
	i := 1

	// 関数の引数は値コピーなので呼び出し元には関係ない
	inc(i)
	inc(i)
	fmt.Println(i) // == 1

	// ポインタを使って変更する
	inc_p(&i)
	inc_p(&i)
	inc_p(&i)
	fmt.Println(i) // == 4

}

func inc(i int) {
	i++
}

func inc_p(i *int) {
	*i++
}
