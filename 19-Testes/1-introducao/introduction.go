package main

import (
	"fmt"
	address "introducao-testes/addres"
)

func main() {
	typeAddress := address.TypeAddress("Avenida Santos Dumont")
	fmt.Println(typeAddress)
}
