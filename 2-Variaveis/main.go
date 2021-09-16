package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	// inferencia de tipo, ou tipo da viaravel definida pelo seu valor
	variavel2 := "variavel 2"
	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)

	const constante string = "constante 1"
	fmt.Println(constante)

	x, y := "x", "y"

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)
}