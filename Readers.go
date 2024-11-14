package main

import (
	"fmt"
	"io"      // Provides basic interfaces for I/O operations (like Reader and Copy).
	"strings" // Contains functions for manipulating strings, like creating a Reader that reads from a string.
)

func main() {
	// A Reader in Go is an object that implements the io.Reader interface, which is used to read a sequence of bytes from a data source.
	r := strings.NewReader("Hello, Reader!") // This creates a new Reader that reads from the string "Hello, Reader!". // So now, r is a reader that reads the string "Hello, Reader!".
	b := make([]byte, 8)                     // This creates a byte slice ([]byte) of length 8. Byte slices are essentially arrays of bytes in Go. make is a built-in function in Go used to allocate memory for slices, maps, and channels.

	// this is an infinite loop and will keep runing until you break out of it
	for {
		n, err := r.Read(b)                               //this is calling the Read method on r. then read the data into the byte slice b. reads "Hello, R". into b
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b) // n will be 8,
		fmt.Printf("b[:n] = %q\n", b[:n])                 // b[:n] will contain "Hello, R".
		if err == io.EOF {
			break
		}
	}

}
