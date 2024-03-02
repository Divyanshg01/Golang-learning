package main

import "fmt"

func main(){

	sum := 1
	// for i:=0 ; i<5 ;i++{
	// 	sum++;
	// }

	//can act as a while like
	// for ;sum<1000;{
	// 	sum +=sum
	// }
	//we can rewrite it as 
	for sum < 1000 {
		sum += sum
	}
	// infinite loop
	// for {

	// }
	fmt.Println(sum)
}