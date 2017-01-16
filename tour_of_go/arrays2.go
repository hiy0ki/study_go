package main

import (
	"fmt"
	"strings"
)

func main() {
	// 多次元配列的なslice
	// create a tic-tac-toe board.
	board := [][]string{
		[]string{"_","_","_",},
		[]string{"_","_","_",},
		[]string{"_","_","_",},
	}


	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append . sliceに要素を追加する
	var s []int
	printSlice(s)

	s = append(s, 1)
	printSlice(s)
	
	s = append(s, 2,3,4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
