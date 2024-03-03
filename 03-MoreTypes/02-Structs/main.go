package main

import "fmt"
type Vertex struct{
	X int
	Y int
}
var (
	v1 = Vertex{1,2}
	v2 = Vertex{X : 1}
	v3 = Vertex{Y : 1}
	p = &Vertex{43,2}
)
func main() {

	// v := Vertex{1,2}
	// v.X = 4

	// p := &v
	// //no arrow operator needed
	// p.Y = 10


	fmt.Println(*p)
	// fmt.Println("Hello")
}