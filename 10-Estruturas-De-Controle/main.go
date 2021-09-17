package main

import "fmt"

func main() {
	number := 10

	if number > 15 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if init

	// a variavel fica limitada a escopo do if, ou seja, fora do escopo ela nao existe
	if OtherNumber := number; OtherNumber > 0 {
		fmt.Println("OtherNumber is greater than zero")
	}
}