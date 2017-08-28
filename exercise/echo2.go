package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep, is := "", "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		is += sep + fmt.Sprint(i)
		sep = " "
	}

	fmt.Println(s)
	fmt.Println(is)
}
