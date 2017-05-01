package main

import (
	"fmt"
)

/*
 配列を渡すとその順番を逆にして表示する
*/
func main() {
	a := [4]string{
		"yamada", 
		"suzuki",
		"tanaka",
		"satou", // この要素にカンマがないとsyntax errorになる
	}

	fmt.Println(a)

	length := len(a)
	fmt.Println(length)

	// その１
	out_ary := []string{}
	for i := 1; i <= length; i++ {
		out_ary = append(out_ary, a[length -i])
	}
	
	fmt.Println(out_ary)

	// その２
	out_ary2 := []string{}
	for i := length -1 ; i >= 0; i-- {
		out_ary2 = append(out_ary2, a[i])
	}
	
	fmt.Println(out_ary2)

}
