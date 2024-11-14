package main

import "fmt"

type I interface {
	MC()
}

type T struct {
	S string
}

func (t T) MC() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello WOrd"}
	i.MC()
}
