package main

import "fmt"

func main() {
	employeeAndSalary := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15325.78,
		"Pedro Junior":   1200.00,
	}

	employeeAndSalary["Rafael Filho"] = 1350.02
	delete(employeeAndSalary, "absent")

	for name, salary := range employeeAndSalary {
		fmt.Println(name, salary)
	}
}
