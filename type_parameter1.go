package main

import "fmt"

type List[T comparable] struct {
	next *List[T]
	val  T
}

// Add new elements in front of the list
// the T in the method signature is a type parameter, which means it's a placeholder for a specific type that will be specified when the List struct is instantiated.
func (l *List[T]) Insert(val T) *List[T] { // l is a pointer to a List struct of type T, l is a receiver name, *List[T] is the receiver type
	return &List[T]{next: l, val: val} // next: a pointer ro an existing list
}

// Append adds a new element at the end of the list.

func (l *List[T]) Append(val T) *List[T] {

	if l == nil {
		// if the list is empty, create a new list
		return &List[T]{val: val}
	}
	// Traverse to the end of the list and add the new value
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: val}
	return l
}

// Find searches for a value in the list and returns the node if found, or nil if not.

func (l *List[T]) Find(val T) bool {
	if l.val == val {
		return true
	}
	if l.next != nil {
		return l.next.Find(val)
	}
	return false
}

// Length returns the number of elements in the list
func (l *List[T]) Length() int {
	length := 0
	current := l
	for current != nil {
		length++
		current = current.next
	}
	return length

}

// Print prints the elements of the list.
func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var list *List[int] // Start with an empty list of ints

	// Insert elements
	list = list.Insert(10)
	list = list.Insert(20)
	list = list.Insert(30)

	// Append elements
	list = list.Append(40)
	list = list.Append(50)

	// Print the list
	fmt.Print("List: ")
	list.Print()

	// Find a value in the list
	valToFind := 20
	found := list.Find(valToFind)
	if found != nil {
		fmt.Printf("Found %d in the list.\n", valToFind)
	} else {
		fmt.Printf("%d not found in the list.\n", valToFind)
	}

}
