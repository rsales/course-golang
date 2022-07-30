package main

import "fmt"

func main() {
	employeeForLetter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15325.78,
			"Guga Pereira":   8045.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(employeeForLetter, "P")

	// example class
	for letter, employee := range employeeForLetter {
		fmt.Println(letter, employee)
	}

	fmt.Println("=============== Divider ==================")

	// test for nested map
	for letter, employee := range employeeForLetter {
		fmt.Println("Letter Name:", letter)
		for emp, salary := range employee {
			fmt.Println(emp, salary)
		}
	}
}
