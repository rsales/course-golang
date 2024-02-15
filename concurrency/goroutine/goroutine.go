package main

import (
	"fmt"
	"time"
)

func speak(people, text string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", people, text, i+1)
	}
}

func main() {
	// speak("Maria", "Pq vc não fala comigo?", 3)
	// speak("joão", "Só posso falar depois de vc!", 1)

	// go speak("Maria", "Ei...", 500)
	// go speak("João", "Opa...", 500)

	go speak("Maria", "Entendi", 10)
	speak("João", "Parabéns", 5)
}
