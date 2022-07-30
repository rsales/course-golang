package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // there is no while in go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Even ")
		} else {
			fmt.Print("Odd ")
		}
	}

	fmt.Println("\nInfinity... ")

	for { // infinity for
		fmt.Println("Forever...")
		time.Sleep(time.Second)
	}
}
