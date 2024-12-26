package main

import (
	"fmt"
	"slices"
)

func main() {
	// what is nil is nil zero valued --> different for other datastructures
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	// zero value for string is the empty string
	// slices do not need to have a length defined -> slices are on top of arrays to work with arrays more easily
	c := make([]string, 10)
	c = append(c, "hey")
	fmt.Println(c, "len:", len(c), "cap:", cap(c))

	some_other := make([]string, 5)
	copy(some_other, c)
	// slicing like in python
	fmt.Println(c[:5])
	fmt.Println(c[2:])

	if slices.Equal(c, some_other) {
		fmt.Println("slices are equal")
	} else {
		fmt.Println("slices are not equal")
	}

	// two dimensional datastructures:
	twoDims := make([][]int, 3)
	fmt.Println(twoDims)

	for i := 0; i < len(twoDims); i++ {
		twoDims[i] = make([]int, 3)
		for j := 0; j < len(twoDims[i]); j++ {
			twoDims[i][j] = 9
		}
	}
	fmt.Println(twoDims)

	fmt.Println(nil)
}
