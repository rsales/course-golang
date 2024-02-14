package main

import (
	"fmt"
	"strings"
)

type people struct {
	name      string
	last_name string
}

func (p people) getFullName() string {
	return p.name + " " + p.last_name
}

func (p *people) setFullName(fullName string) {
	pices := strings.Split(fullName, " ")
	p.name = pices[0]
	p.last_name = pices[1]
}

func main() {
	p1 := people{"Rafael", "Sales"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Pedro Silva")
	fmt.Println(p1.getFullName())
}
