package main

import "fmt"

type curse struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	type dynamic interface{}

	var thing2 dynamic = "Opa"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = curse{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(thing2)
}
