package main

import "fmt"

func main() {
	// var approved map[int]string
	// maps must be initialized
	approved := make(map[int]string)

	approved[89253518030] = "Maria"
	approved[35244261070] = "Pedro"
	approved[16948254004] = "Ana"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[35244261070]) // result "Ana"
	delete(approved, 35244261070)      //remove item map
	fmt.Println(approved)

}
