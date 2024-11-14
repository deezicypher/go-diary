package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[0][2] = "X"
	board[1][1] = "O"
	board[2][0] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}

	nSlice()
}

func nSlice() {
	var s []int
	printS(s)

	s = append(s, 0)
	printS(s)

	s = append(s, 1)
	printS(s)

	s = append(s, 2, 3, 4)
	printS(s)

	n := make([]int, 5, 10)
	printS(n)

	w := make(map[string]int)
	w["apple"] = 1
	w["orange"] = 2

	fmt.Println(w)
}

func printS(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
