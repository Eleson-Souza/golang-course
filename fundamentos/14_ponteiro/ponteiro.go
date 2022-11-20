package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variável "i" e atribuindo ao ponteiro "p"
	*p++   // valor associado ao ponteiro + 1
	i += 2

	// Como 'p' e 'i' apontam para o mesmo endereço de memória, qualquer alteração no valor deles, interferirá no outro, pois esse valor está associado ao endereço que ambos estão apontando.
	fmt.Println(p, *p)
	fmt.Println(&i, i)
}
