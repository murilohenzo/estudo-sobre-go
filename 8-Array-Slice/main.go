package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array [5]int

	array[0] = 1
	array[0] = 2
	array[0] = 3
	array[0] = 4
	array[0] = 5

	fmt.Println(array)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array))
	fmt.Println(reflect.TypeOf(slice))

	slice2 := array2[1: 3]
	fmt.Println(slice2)

	//arrays internos

	slice3 := make([]int, 10, 15)
	fmt.Println("Slice3 Original: ", slice3)
	fmt.Println("Tamanho:", len(slice3)) // length
	fmt.Println("Capacidade:", cap(slice3)) //capacity

	for i := 0; i < 5; i++ {
		slice3 = append(slice3, i)
	}

	fmt.Println("Novo slice3: ", slice3)
	fmt.Println("Tamanho:", len(slice3)) // length
	fmt.Println("Capacidade:", cap(slice3)) //capacity

// estourando capacidade para o go expandir o valor da capacidade
	for i := 0; i < 5; i++ {
		slice3 = append(slice3, i)
	}

	fmt.Println("Estourando a capacidade slice3: ", slice3)
	fmt.Println("Tamanho:", len(slice3)) // length
	fmt.Println("Capacidade:", cap(slice3)) //capacity
}