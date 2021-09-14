package main

import (
	"errors"
	"fmt"
)

func main() {
	// int64 is the set of all signed 64-bit integers.
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	// uint32 is the set of all unsigned 32-bit integers
	var numero2 uint32 = 1000000000
	fmt.Println(numero2)

	// rune is an alias for int32 and is equivalent to int32 
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	// used, by convention, to distinguish byte values from 8-bit unsigned
	var numero4 byte = 123
	fmt.Println(numero4)

	var numero_float1 float32 = 123000000000000000000000000000000000000.45
	fmt.Println(numero_float1)

	var numero_float2 float64 = 12300000000000000000000000000000000000000000000000.45
	fmt.Println(numero_float2)

	numero_float3 := 12300000000000000000000000000000000000000000000000000000000000.454
	fmt.Println("\nPega o tipo de dado via inferencia, pegando o valor float da arquitetura do computador\n", numero_float3)


	var str string = "variavel"
	fmt.Println(str)

	str1 := "variavel 2"
	fmt.Println(str1)

	char1 := 'A'
	fmt.Println("Posicao do caractere na tabela asci\n", char1)

	//valores iniciais em go, numeros comecam no zero e string comeca como vazio, para erros nil que eh nulo no golang

	var text string
	fmt.Println(len(text))

	var number int8
	fmt.Println(number)

	var number_float float32
	fmt.Println(number_float)

	var boolean bool
	fmt.Println(boolean)

	var erro error
	fmt.Println(erro)

	// create error

	var httpError error = errors.New("Internal server error")
	fmt.Println(httpError)
}