package main

import (
	"fmt"
	"time"
)

func validateType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "I don't know"
	}
}

func main() {
	fmt.Println(validateType(9.8))
	fmt.Println(validateType(10))
	fmt.Println(validateType("Hello"))
	fmt.Println(validateType(func() {}))
	fmt.Println(validateType(time.Now()))
}
