package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferĂȘncia de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2
	var w, z byte = 1, 2
	fmt.Println(x, y)
	fmt.Println(w, z)

	x, y = y, x
	fmt.Println(x, y)
}
