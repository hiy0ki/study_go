package main

import (
	"fmt"
	"os"
)

func main() {
	// env set
	os.Setenv("HOGE", "hogehoge")

	// env get
	hoge := os.Getenv("HOGE")

	// == "hogehoge"
	fmt.Println(hoge)

	// show env list
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
