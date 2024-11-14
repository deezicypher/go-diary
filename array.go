package main

import "fmt"

func main() {
	var a [10]string
	var b [2]string
	a[0] = "Hey"
	b[0] = "you"
	a[1] = "I'm talking to you"
	fmt.Println(a[0], b[0])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	s := primes[1:3]
	fmt.Println(s)
	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{4, false},
		{6, true},
	}
	fmt.Println(st)
}
