package main

import "fmt"

func main() {
	const pi float64 = 3.14159265358979323846
	const pi2 = 3.14
	// fmt.Println(pi)
	// fmt.Println(pi2)
	fmt.Println("p1:", pi)
	fmt.Println("p2:", pi2)

	// variables enteras
	base := 12
	var height int = 10
	var area int

	fmt.Println(base, height, area)

	// zero value
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// area cuadrado
	const baseCuadrado = 10
	areaCua := baseCuadrado * baseCuadrado

	fmt.Println("area cuadrado:", areaCua)
}
