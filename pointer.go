package main

import "fmt"

//A pointer holds the memory address of a value.

func Pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func main() {
	i := 1
	p := &i       //Point to i,this assigns the address of i to p (using the & operator to get the address of i).
	fmt.Print(p)  // prints the memory address that p is pointing to
	fmt.Print(*p) // Dereferencing p: prints the value stored at the address, which is 42

}
