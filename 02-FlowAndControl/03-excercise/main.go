package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	fmt.Printf("%0.3f" ,z)
	for i:=1 ;i<10;i++ {
		if z == x {
			break
		}else {
			z -= (z*z -x) / (2*z)
			// fmt.Println(z)
	fmt.Printf("%0.3f" ,z)


		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}