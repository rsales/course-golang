package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for index, number := range numbers {
		fmt.Printf("%d) %d\n", index, number)
	}

	for _, number := range numbers {
		fmt.Println(number)
	}
}
