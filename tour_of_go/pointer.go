package main

import "fmt"

func main() {
	i, j := 42,2701

	p := &i // & はそのオペランドへのポインタを引き出す
	fmt.Println(*p) // * はポインタの指す先の変数をしめす
	*p = 21
	fmt.Println(i)
	
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
