package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := randomNumber(); i > 5 { // also supported on switch
		fmt.Println("You Won!")
		fmt.Print(i)
	} else {
		fmt.Println("You Lost!")
		fmt.Print(i)
	}

	// fmt.Print(i) // Code Error
}
