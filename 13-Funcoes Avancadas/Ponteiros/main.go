package main

import "fmt"

func reverse(number int) int {
	return number * -1
}

// nao precisa de retorno, pois estamos passando a referencia
// mudando o valor no local de memoria da variavel number
func reverseWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20
	reverseNumber := reverse(number)
	fmt.Println(reverseNumber)
	fmt.Println(number)

	number2 := 40
	fmt.Println(number2)
	reverseWithPointer(&number2)
	fmt.Println(number2)
}