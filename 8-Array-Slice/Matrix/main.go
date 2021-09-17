package main

import "fmt"

func main() {
	// cria uma matriz de 5 posicoes vazias
	matrix := make([][]uint8, 5)
	fmt.Println(matrix)
	for i := range matrix {
		// popula as 5 posicoes da matriz com 0
		matrix[i] = make([]uint8, 5)
	}
	fmt.Println(matrix)
}