package main

import "fmt"

func main() {

	// ciclo for
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	// for while
	counter := 0

	fmt.Printf("\n")
	for counter < 10 {
		fmt.Printf("%d ", counter)
		counter++
	}

	// for forever
	counterForever := 0
	fmt.Printf("\n")

	for {
		fmt.Printf("%d ", counterForever)
		counterForever++

		if counterForever == 10 {
			break
		}
	}
}
