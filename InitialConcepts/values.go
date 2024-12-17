package main

import "fmt"

func main(){
	// strings can be concatenated
	fmt.Println("Hello" + "World")
	// PrintLn automatically converts input values to string reprresentation
	fmt.Println("1+1" , 1+1)
	// floats
	fmt.Println("7/3 = " ,  7/3)

	fmt.Println(true && true)

	fmt.Println(!true)

	fmt.Println(false || false)


}