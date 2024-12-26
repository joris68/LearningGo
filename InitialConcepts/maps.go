package main

import (
	"fmt"
	"maps"
)

func main() {

	myFirstMap := make(map[string]int)
	myFirstMap["five"] = 5
	myFirstMap["six"] = 6
	fmt.Println(myFirstMap)

	mySecondMap := make(map[string]int)
	mySecondMap["five"] = 5
	mySecondMap["six"] = 6
	fmt.Println(mySecondMap)
	delete(mySecondMap, "seven")
	fmt.Println(mySecondMap)
	clear(mySecondMap)
	fmt.Println(mySecondMap)
	fmt.Println(myFirstMap["bla"])

	fmt.Println(maps.Equal(myFirstMap, mySecondMap))

	// maps retunr zero -value of the type if key not found : important to use this concept !!!!
	value, ok := myFirstMap["five"]
	fmt.Println(value, ok)
}
