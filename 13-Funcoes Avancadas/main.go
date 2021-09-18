package main

import "fmt"

// funcao com retorno nomeado
func mathematicalCalculations(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

// funcao variatica, nao precisa especificar a quantidade de parametros a serem passados

func mathematicalCalculations2(numbers ...int) (sum int) {
	aux := 0
	for _, number := range numbers {
		aux += number
	}
	sum = aux
	return
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

// funcao recursiva
func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position - 1)
}
 
func main() {
	fmt.Println(mathematicalCalculations(2, 3))
	fmt.Println(mathematicalCalculations2(1, 2, 3, 4, 5, 6))
	write("Hello, world!", 1, 3, 4, 5, 5)

	// funcoes anonimas 
	func() {
		fmt.Println("funcao anonima")
	}()

	func(text string) {
		fmt.Println(text)
	}("Funcao anonima com parametros")
	
	retornoDeUmaFuncaoAnonima := func(text string) string {
		return fmt.Sprintf("Recebido -> %s %d", text, 10)
	}("Funcao anonima com parametros")

	fmt.Println(retornoDeUmaFuncaoAnonima)

	for i := uint(0); i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}