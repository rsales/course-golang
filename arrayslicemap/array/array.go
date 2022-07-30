package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notes [3]float64
	fmt.Println(notes)

	notes[0], notes[1], notes[2] = 7.8, 4.3, 9.1
	fmt.Println(notes)

	total := 0.0
	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	media := total / float64(len(notes))
	fmt.Printf("Média %.2f\n", media)
}
