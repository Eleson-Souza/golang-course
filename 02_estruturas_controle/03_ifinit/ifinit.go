package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// variável criada e visível apenas no escopo do if/else
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!!!")
	}
}
