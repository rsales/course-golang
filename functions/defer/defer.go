package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("End!")
	if number > 5000 {
		fmt.Println("High Value...")
		return 5000
	}
	fmt.Println("Low Value...")
	return number
}

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(3000))
}
