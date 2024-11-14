package main

import (
	"fmt"
)

func main() {
	sum := 1
	for sum <= 5 {
		sum += sum
		fmt.Println(sum)
	}

}

func loop() {
	sum := 0
	for i := 0; i <= 5; i++ {
		sum += sum
		fmt.Println(sum)
	}
}
