package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix (pos fixada)
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)

	// error
	//fmt.Println(x == y--)
}
