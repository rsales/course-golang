package main

import "fmt"

func main() {
	var b byte = 3                   // Convencional
	fmt.Println("[Conventional]", b) // result: 3

	i := 3                           // inferência de tipo
	fmt.Println("[Include type]", i) // result: 3
	i += 3                           // i = i + 3
	fmt.Println("[+=]", i)           // result: 6
	i -= 3                           // i = i - 3
	fmt.Println("[-=]", i)           // result: 3
	i /= 2                           // i = i / 2
	fmt.Println("[/=]", i)           // result: 1
	i *= 2                           // i = i * 2
	fmt.Println("[*=]", i)           // result: 2
	i %= 2                           // i = i % 2
	fmt.Println("[%=]", i)           // result: 0

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
