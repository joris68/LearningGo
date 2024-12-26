package main

import "fmt"

func FibSeq(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return FibSeq(n-1) + FibSeq(n-2)
}

// multiple return values in Go
func multiple() (int, string) {
	return 1, "hey"
}

// go does not have a return type like void, if nothing is returned then nothing will get returned
func NoReturnType() {
	fmt.Println("this function has no return type")
}

func foo(a ...int) int {
	result := 0
	for _, val := range a {
		result += val
	}
	return result
}

func some() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}

// anonymous funtion

func main() {

	result := FibSeq(12)
	number, _ := multiple()
	fmt.Println(result)
	fmt.Println(number)
	var myArray = []int{1, 2, 4}
	fmt.Println(foo(myArray...))

	// higher order functions
	a := some()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	//
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

}
