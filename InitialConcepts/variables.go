package main

import "fmt"

func main(){

	// you can do type annotation
	var a string = "hello"
	var b string = "world"

	// you can do it without type annonation
	// the compiler can infer type automatically -->
	var d = "hey"
	var e = "there"

	//you can define several variables in a row
	//var x , y, z int = 1 ,"boo" ,3 -> not possible
	// var x int , y int, z int = 1 ,2 ,3 will throw syntax error
	var x , y, z int = 1 ,2 ,3 // here now all must be strings

	var oki bool = true

	// declaring without actually initializing it -> will be zero initilaized
	var some_int int
	var some_bool bool
	var some_string string
	var some_float float32


	fmt.Println(a + b)
	fmt.Println(d + e)
	fmt.Println(x , y, z)
	fmt.Println(oki)

	fmt.Println(some_int) //zero initalized with 0
	fmt.Println(some_bool) // zero init with false
	fmt.Println(some_string) // zero init with empty string
	fmt.Println(some_float) // zer init with 0

	// short form for decalring and initalizing a variable
	f := "hey"
	fmt.Println(f)
}