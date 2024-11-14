package main

import (
	"fmt"
	"math"
)

type I interface {
	MC()
}

type T struct {
	S string
}

func (t *T) MC() {
	fmt.Println(t.S)
}

type F float64

func (f F) MC() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.MC()
	i = F(math.Pi)
	describe(i)
	i.MC()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
