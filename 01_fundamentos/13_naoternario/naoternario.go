package main

import "fmt"

// Não tem operador ternário :(
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println("Aluno 1:", obterResultado(8))
	fmt.Println("Aluno 2:", obterResultado(5.9))
}
