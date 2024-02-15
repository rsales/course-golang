package main

import "fmt"

type sporty interface {
	turnOnTurbo()
}

type luxurious interface {
	makeParallelParking()
}

type sportyLuxurious interface {
	sporty
	luxurious
	// create more methods
}

type bwm7 struct{}

func (b bwm7) turnOnTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) makeParallelParking() {
	fmt.Println("Baliza...")
}

func main() {
	var b sportyLuxurious = bwm7{}
	b.turnOnTurbo()
	b.makeParallelParking()
}
