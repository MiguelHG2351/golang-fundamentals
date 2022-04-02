package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 20

	// sumar
	result := x + y
	fmt.Println("Sumar: ", result)

	// restar
	result = x - y
	fmt.Println("Restar: ", result)

	// multiplicar
	result = x * y
	fmt.Println("Multiplicar: ", result)

	// dividir
	result = x / y
	fmt.Println("Dividir: ", result)

	// modulo
	result = x % y
	fmt.Println("Modulo: ", result)

	// incremento
	x++
	fmt.Println("Incremento: ", x)

	// decremento
	x--
	fmt.Println("Decremento: ", x)

	// exponencial
	result = x ^ y
	fmt.Println("Exponencial: ", result)
}
