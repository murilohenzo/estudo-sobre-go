package main

import "fmt"

func closure() func() {
	text := "closure function"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	result := closure()
	result()
}