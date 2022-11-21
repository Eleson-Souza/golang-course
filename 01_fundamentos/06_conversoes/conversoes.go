package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println("x é do tipo", reflect.TypeOf(x))
	fmt.Println("y é do tipo", reflect.TypeOf(y))
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado... 123 = unicode "{"
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste " + strconv.Itoa((123)))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
