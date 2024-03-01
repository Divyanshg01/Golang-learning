package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

//naked return
func splitup(sum int)(x , y int){
	x = sum * 4/9
	y = sum -x
	return
}
func main() {
	a, b := swap("Hello", "World")
	l ,m := splitup(21)
	fmt.Println(a,b)

	fmt.Println(splitup(21))
	fmt.Println(l ,m)
}