package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./hoge.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(string(buf[:n]))
	}

}
