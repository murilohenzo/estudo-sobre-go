package main

import "fmt"

func handle1() {
	fmt.Println("exec handle 1")
}

func handle2() {
	fmt.Println("exec handle 2")
}

func main() {

	defer handle1()
	// defer = Adiar a execucao de um pedaco de codigo
	handle2()

}