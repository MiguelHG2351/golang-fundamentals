package main

import (
	"fmt"
	"log"
	"strconv"
)

func isPair(i int) bool {
	return i%2 == 0
}

func main() {
	var valor int = 1

	if valor == 1 {
		fmt.Printf("Valor igual a 1\n")
	} else {
		fmt.Printf("Valor diferente de 1\n")
	}

	// tiene los mismo operadores que otros lenguajes, que pereza agregar más ejemplos xd

	value, error := strconv.Atoi("1")
	// value, error := strconv.Atoi("sus")

	if error == nil {
		fmt.Printf("Valor convertido a entero: %d\n", value)
	} else {
		log.Fatal("Value:", error)
	}
}
