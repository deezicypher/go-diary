package main

import (
	"fmt"
	"time"
)

// The error type is a built-in interface similar to fmt.Stringer:
// type error interface {
//    	Error() string
//  }

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &MyError{
		time.Now(),
		"Didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
