package main

import "fmt"

type car struct {
	name        string
	actualSpeed int
}

type ferrari struct {
	car     // anonymus fields
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.actualSpeed = 0
	f.turboOn = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.name, f.turboOn)
	fmt.Println(f)
}
