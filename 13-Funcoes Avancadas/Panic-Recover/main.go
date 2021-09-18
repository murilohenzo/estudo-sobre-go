package main

import "fmt"

func recoverExec() {
	if r:= recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execucao")
	}
}

func mean(n1, n2 float64) bool {
	defer recoverExec()
	mean := (n1 + n2) / 2

	if mean > 6 {
		return true
	} else if mean < 6 {
		return false
	}
	// vai interromper o fluxo do programa
	panic("mean is 6")
}

func main() {
	fmt.Println(mean(6, 9))
	fmt.Println(mean(6, 6))
}