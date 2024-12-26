package main

//import "fmt"

func main() {
	// arrays in go are typed by the number and the type they contain
	// declaration --> will be zero initialized
	var array [5]int

	//fmt.Println(array) // [0,0,0,0,0]

	array[4] = 100

	//fmt.Println(array)

	// len is a built in array function in go
	//fmt.Println(len(array))

	// decalring and initilaizing in one line
	// arrays must have the length defined and declaration --> allocates memory
	b := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(b)

	c := [...]int{1, 3: 10, 12}
	//fmt.Println(c)

	var twoDim [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			twoDim[i][j] = i + j
		}
	}

	//fmt.Println(twoDim)

	// or do it without the for loop

	var anotherTwoDim = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	//fmt.Println(anotherTwoDim)

	var some bool = true
	var some_other int = 12

	//fmt.Println(some, some_other)
}
