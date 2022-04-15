package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	nombre := "Juan"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)

	fmt.Printf(message)

	// tipo de datos
	fmt.Printf("Hello message: %T\n", message)
	fmt.Printf("Cursos: %T\n", cursos)
}
