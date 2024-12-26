package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// command line arguments in go
	fmt.Println(os.Args[0]) // first one is the path of the exececutable
	fmt.Println(os.Args[1]) // that is what you pass in
	if len(os.Args) != 2 {
		log.Fatalln("Not enough arguments specified")
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Inavlid address")
	} else {
		fmt.Printf("the address is %s", addr.String())
	}
}
