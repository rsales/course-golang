package main

import "fmt"

type sport interface {
	turnOnTurbo()
}

type ferrari struct {
	model        string
	tubo_on      bool
	actual_speed int
}

func (f *ferrari) turnOnTurbo() {
	f.tubo_on = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnTurbo()

	var car2 sport = &ferrari{"F40", false, 0}
	car2.turnOnTurbo()

	fmt.Println(car1, car2)
}
