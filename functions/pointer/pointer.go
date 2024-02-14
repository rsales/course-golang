package main

import "fmt"

func inc1(n int) {
	n++
}

// Attention: inc2() It is no longer a pure function
// review: a poointer is representative for a *
func inc2(n *int) {
	// review: * used for dereferencing, in other words
	// have access to value at which the pointer points
	*n++
}

func main() {
	n := 1

	inc1(n) // for value
	fmt.Println(n)

	// review: & used for to obtain the variable address
	inc2(&n) // for reference
	fmt.Println(n)
}
