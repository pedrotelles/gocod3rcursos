package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// método: funcao com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
func main() {
	var p1 produto
	p1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(p1, p1.precoComDesconto())
	p2 := produto{"Televisão", 1989.90, 0.10}
	fmt.Println(p2, p2.precoComDesconto())
}
