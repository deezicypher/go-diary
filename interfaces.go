package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser               //a is declared as a variable of type Abser, meaning it can hold any value that implements the Abser interface.
	f := MyFloat(-math.Sqrt2) //creates a MyFloat variable f. This is a custom type based on float64. The Abs() method is implemented for MyFloat, which means f implements the Abser interface.
	v := Vertex{3, 4}         //creates a Vertex struct. In the next step, a = &v assigns a pointer to v (&v) to the Abser variable a. This works because the Abs() method is defined for the pointer type *Vertex, so *Vertex implements Abser.

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}
