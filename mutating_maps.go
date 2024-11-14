package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42

	fmt.Println("The answer is:", m["Answer"])

	ans := m["Answer"]
	fmt.Println("The answer is:", ans)

	delete(m, "Answer")
	fmt.Println("The answer is:", ans)

	ans, present := m["Answer"]
	fmt.Println("The answer is:", ans, "Present?", present)

}
