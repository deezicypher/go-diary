package main

import "fmt"

var powx = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	for i, v := range powx {
		fmt.Printf("i =%d , v = %d \n", i, v)
	}
}
func rangeOmit() {
	pow := []int{2, 4, 6}

	for _, value := range pow {
		fmt.Printf("Omiting the index, v =%v", value)
	}

	for i := range pow {
		fmt.Printf("Omiting the value, i = %d", i)
	}
}
