package main

import "fmt"

func printApproved(approved ...string) {
	fmt.Println("List of approved")
	for i, approved := range approved {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	approved := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	printApproved(approved...)
}
