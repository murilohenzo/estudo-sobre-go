package main

import "fmt"

func main() {
	var value1 int = 1

	var value2 int = value1

	fmt.Println(value1, value2)


	value1++

	fmt.Println(value1, value2)

	//Ponteiro referencia de memoria
	var value3 int = 100
	var pointer *int = &value3

	fmt.Println(value3, pointer)

	value3 = 150

	fmt.Println(value3, pointer)

	fmt.Println(value3, *pointer) //desreferenciacao

}