package main

import "fmt"

func noteForConcept(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "C"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(noteForConcept(9.8))
	fmt.Println(noteForConcept(6.9))
	fmt.Println(noteForConcept(2.1))
}
