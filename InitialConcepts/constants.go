package main

import (
	"fmt"
	"math"
)

const constant string = "I am a constant"

func main(){
	// I can do a constant expression

	const expr = 5000 / 30

	// explicit type conversion here
	fmt.Println(int64(expr))
	// there is probably implicit type conversion
	// when calling the math.sin method
	fmt.Println(math.Sin(expr))
}