package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	// "_": descarta o elemento
	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
