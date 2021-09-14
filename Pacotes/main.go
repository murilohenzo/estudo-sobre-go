package main

import (
	"fmt"
	"module/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Starting golang")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("johndoe@gmail.com")
	fmt.Println(erro)
}