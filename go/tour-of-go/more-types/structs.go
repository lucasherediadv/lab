package main

import "fmt"

// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // structs fields are accessed using a dot
	fmt.Println(v.X)
}
