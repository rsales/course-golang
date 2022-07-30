package main

import "fmt"

// There are no ternary operators in Go
func getResult(note float64) string {
	if note >= 6 {
		return "Approved"
	}
	return "Disapproved"

	// Ternary notation 😢
	// return note >= 6 ? "Approved" : "Disapproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
