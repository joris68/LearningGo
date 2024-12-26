package main

import "fmt"

type person struct {
	age  int
	name string
}

func personGenerator(age int, name string) *person {
	newPers := person{age, name}
	return &newPers
}

func (p person) someMethod() {
	fmt.Println(p.name)
}

func main() {
	hannes := personGenerator(19, "joris")

	// dereferencing
	fmt.Println(&hannes)
	fmt.Println(hannes)
	fmt.Println(*hannes)
	hannes.someMethod()

}
