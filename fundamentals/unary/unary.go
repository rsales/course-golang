package main

import "fmt"

func main() {
	x := 1
	y := 2

	// ==============
	// only postfix
	// ==============

	x++                           // x += 1 ou  x = x + 1
	fmt.Println("Additive", x)    // result: 2
	y--                           // x -= 1 ou  x = x - 1
	fmt.Println("Subtractive", y) // result: 1

	// fmt.Println(x == y--) // error
}
