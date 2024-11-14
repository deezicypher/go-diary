package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abst() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scalex(f float64) { // The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
	v.X = v.X * f
	v.Y = v.Y * f
}

//Rewritten Functions

func ScaleF(v *Vertex, f float64) { //The function Scale takes a pointer to a Vertex as its first argument: func Scale(v *Vertex, f float64).
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scalex(10)
	ScaleF(&v, 10) //By passing the address of v, you allow the Scale function to directly modify the original Vertex struct. If you passed v by value (without the &), a copy of v would be passed to Scale, and the changes made inside the function would not affect the original
	fmt.Println(v.Abs())

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleF(p, 3)
	fmt.Println(p)

}
