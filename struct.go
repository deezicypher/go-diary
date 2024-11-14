package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"obi", 24}
	p.age = 26
	pt := &p
	fmt.Printf("%p \n", pt)
}
