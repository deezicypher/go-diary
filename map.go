package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // make is a built-in function in Go used to allocate memory for slices, maps, and channels.
	m["Bell Labs"] = Vertex{
		50.678, -70.8989954,
	}
	var ml = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m["Bell Labs"], "\n", ml)
}
