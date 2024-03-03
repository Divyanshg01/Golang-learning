package main

import "fmt"

func main() {
	i , j:= 10 ,220
	var q *int = &j
	p := &i
	fmt.Println(i , j ,*p , p ,*q, &i ,q ,&j)
	//go has no pointer arithmatic
	*p +=2
	fmt.Println(*p)
}