
package main

import "fmt"

func main(){
	// for loop with single condition
	//i := 1
	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// a more classical one
	for j := 0; j < 3; j++ {
		fmt.Println("doing something in the loop")
	}

	// using range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// there are no while loops in go -> so we can use for with no condition
	for {
		fmt.Println("here should be a break statement")
		break
		continue // we can also use continue in a for loop
	}

	// also we


}