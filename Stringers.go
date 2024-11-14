package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years))", p.Name, p.age)
	// fmt.Print - Directly prints to the console,	No return value (just writes to output)
	// fmt.Sprint -	Returns a formatted string,	Returns a string (does not print)
}

func main() {
	a := Person{"Zed", 23}
	d := Person{"Dan", "34"}
	fmt.Println(a, d)
}
