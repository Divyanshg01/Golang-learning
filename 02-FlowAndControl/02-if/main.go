package main

import (
	"fmt"
	"math"
)
func sqrt(x float64) string {
	if x<0 {
		return sqrt(-x) +"i"
	}
	return fmt.Sprint(math.Sqrt(x))

	//Sprintf returns the result in the form of string


}



func main(){
	fmt.Println(sqrt(2) , sqrt(16))
	if v:=19 ; 2>3 {

		fmt.Println("inside if" ,v)
	}
	//here v can not be accessed outside of if block but are accessible in else too
}