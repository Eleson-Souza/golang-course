package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678912] = "Maria"
	aprovados[12345678934] = "Pedro"
	aprovados[12345678956] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678956])
	delete(aprovados, 12345678956)
	fmt.Println(aprovados[12345678956])
}
