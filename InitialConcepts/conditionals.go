package main

import "fmt"

func main(){

	var i int = 5

	if i > 4 {
		fmt.Println("5 bigger than 4")
	}

	if i % 2 == 0 || i % 2 == 1 {
		fmt.Println("I am always true")
	}

	// js can now be used in every other branch in these conditionals
	if j := 8; j < 10 {
		fmt.Println("j smaller 10")
	} else if j < 11 {
		fmt.Println("j smaller 11")
	} else if j < 12 {
		fmt.Println("j smaller 12")
	} else{
		fmt.Println("j is too big amk")
	}

}