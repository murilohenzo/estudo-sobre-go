package main

import (
	"fmt"
	"math"
)

type RetanguloDto struct {
	altura  float64
	largura float64
}

func (r RetanguloDto) area() float64 {
	return r.altura * r.largura
}

type CirculoDto struct {
	raio float64
}

func (c CirculoDto) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type IForm interface {
	area() float64
}

// os structs vao precisar criar essa assinatura d0 metodo area(), para fazer o uso dessa funcao
func escreverArea(form IForm) {
	fmt.Printf("A area da forma eh %0.2f", form.area())
}

func main() {
	// implementação implicita da interface
	r := RetanguloDto{10, 15}
	escreverArea(r)

	c := CirculoDto{10}
	escreverArea(c)

}