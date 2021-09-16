package main

import "fmt"

type Person struct {
	name string
	lastName string
	age uint8
	height uint8
}

type Student struct {
	Person
	course string
	university string
}

func main() {
	p1 := Person{"John", "Doe", 22, 180}
	fmt.Println(p1)

	e1 := Student{p1, "Engenharia", "USP"}
	fmt.Println(e1)

}