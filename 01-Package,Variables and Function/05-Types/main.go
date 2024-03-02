package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	fmt.Println(needInt(Small))
	
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needFloat(Big)): This line calls the needFloat function with the Big constant (also an integer) as the argument. Similar to the previous line, Big is implicitly converted to a float64. However, because Big is a very large number, the conversion may not be accurate due to limitations in floating-point representation. The printed result might be inaccurate or an approximation depending on the system architecture.

	// fmt.Println(needInt(Big))
}
// import (
	// "fmt"
	// "math"
	// "math/cmplx"
	// "math/cmplx"
// )

//Some different datatypes
// byte // alias for uint8

//rune // alias for int32
// represents a Unicode code point

//float32 float64

//complex64 complex128
// const (
// 	Big = 1<<100
// 	Small  = Big >> 99
// )

// func main(){
// 	// not :=
// 	const a = 12
// 	const b = true
// 	var x ,y int = 3 ,4
// 	f   := math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)
// 	fmt.Println(x,y,z ,a,b)

	// var (
	// 	ToBe bool = false
	// 	MaxInt uint64 = 1<<64 -1
	// 	z complex128 = cmplx.Sqrt(-5 + 12i)
	// )

	// // fmt.Println("Hello \n")
	// fmt.Printf("%T : %v \n", ToBe , ToBe)
	// fmt.Printf("%T : %v \n" ,MaxInt ,MaxInt)
	// fmt.Printf("%T : %v" , z ,z)
// }

// int is 32bit on 32 bit system and 64 on 64 bit system
//Default values like var i init
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.


//when not explicitly defined the new variable may be an int, float64, or complex128 

