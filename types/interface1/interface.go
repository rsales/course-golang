package main

import "fmt"

type printable interface {
	toString() string
}

type people struct {
	name      string
	last_name string
}

type product struct {
	name  string
	price float64
}

// interfaces they are implemented implicity
func (p people) toString() string {
	return p.name + " " + p.last_name
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func printer(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = people{"Roberto", "Silva"}
	fmt.Println(thing.toString())
	printer(thing)

	thing = product{"Jeans", 79.90}
	fmt.Println(thing.toString())
	printer(thing)

	p2 := product{"Jeans", 179.90}
	printer(p2)
}
