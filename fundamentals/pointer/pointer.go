package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // create pointer

	p = &i                       // Action: getting the address of variable 'i' in memory
	println("Value of 'p':", *p) // return: 1
	println("Value of 'i':", i)  // return: 1
	*p++                         // Action: dereferencing (taking the value) and additive value
	println("Value of 'p':", *p) // return: 2
	println("Value of 'i':", i)  // return: 2
	i++                          // Action: Additive value
	println("Value of 'p':", *p) // return: 3
	println("Value of 'i':", i)  // return: 3

	// Go doesn't have pointer arithmetic
	// p++ // return error in code

	fmt.Printf("Memory Address 'p': %v, Value of 'p' [Dereferencing]: %v, Value of 'i': %v, Memory Address 'i': %v", p, *p, i, &i)
}
