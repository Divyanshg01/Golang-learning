package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	t := time.Now()
		switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Print("Go runs")
	switch os:=runtime.GOOS;os{
	case "darwin" :
		fmt.Println("OS X.")
	case "linux" :
		fmt.Println("linux")
	default : 
		fmt.Printf("%s " ,os)		
	}
	
}

// Switch without a condition is the same as switch true.

// This construct can be a clean way to write long if-then-else chains.