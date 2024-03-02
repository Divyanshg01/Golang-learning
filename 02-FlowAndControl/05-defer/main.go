package main

import "fmt"
func again(){
	 fmt.Println("done")

}
func main() {
	// fmt.Println("Hello world 1")
	// defer fmt.Println("Hello world from Defer 2")
	// fmt.Println("Hello world 3")
	// fmt.Println("Hello world 4 ")
	// fmt.Println("Hello world 5")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	again()
}

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.