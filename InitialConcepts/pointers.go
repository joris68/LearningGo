package main

import "fmt"

func zeroval(ival int) {

}

func zeropt(iptr *int) {

}

func main() {

	x := 7
	// here I reference x the varibale x
	y := &x
	fmt.Println(x)
	fmt.Println(y)

	// dereferencing, y is a pointer here, here you actually get the value from the pointer y be dereferencing
	z := *y
	fmt.Println(z)

	// this returns a pointer
	w := new(int)
	fmt.Println(w)
	// ths is dereferencing --> setting the value
	*w = 10
	fmt.Println(w, *w)

}
