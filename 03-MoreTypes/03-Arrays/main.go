package main

import "fmt"

func main() {
	// var a[n]int
	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "world"

	// natural := [5]int{1,2,3,4,5}
	// primes := [6]int{2,3,5,7,11,13}

	// fmt.Println(a[0],a[1])
	// fmt.Println(primes)
	// fmt.Println(natural)

	//Slices

	// var s[] int = natural[0:3]//elements from 0th index to 2nd index
	// fmt.Println(s)

	// A slice does not store any data, it just describes a section of an underlying array.

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	//slice literal
	// q := []bool{true , false ,true}
	// fmt.Println(q)
	// sa:= []struct{
	// 	i int
	// 	b bool
	// }{
	// 	{2,true},
	// 	{3,false},
	// }

	//slice default
	// n := [10]int{1,2,3,4,5,6,7,8,9,10}
	// var sab[] int =n[0:]
	// fmt.Println(sa ,sab)
	// 	s := []int{2, 3, 5, 7, 11, 13}

	// 	s = s[1:4]
	// 	fmt.Println(s)

	// 	s = s[:2]
	// 	fmt.Println(s)

	// 	s = s[1:]
	// 	fmt.Println(s)

	//length and capacity of slice
	//length - no. of elements in the slice
	//capacity - no. of elements from the first index of slice till the last index in the orignal array
	// s := []int {2,3,4,5,6}
	// var y []int
	// //default value of y is nill
	// if y == nil {
	// 	fmt.Println("nil")
	// }
	// s = s[2:4]
	// fmt.Println(s , len(s) , cap(s))
	// fmt.Println(y , len(y) , cap(y))
	a := make([]int, 5)// len ->5
	b := make([]int , 3,6) // len -> 3 , cap ->6
	
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	var s[] int
	s = append(s, 2, 3, 4)
	
	fmt.Println(a ,b ,board)
	
}
