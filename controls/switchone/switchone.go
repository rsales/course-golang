package main

import "fmt"

func noteForConcept(n float64) string {
	var note = int(n)
	switch note {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Ivalid note"
	}
}

func main() {
	fmt.Println(noteForConcept(9.8))
	fmt.Println(noteForConcept(6.5))
	fmt.Println(noteForConcept(2.1))
	fmt.Println(noteForConcept(11))
	fmt.Println(noteForConcept(10.9))
}
