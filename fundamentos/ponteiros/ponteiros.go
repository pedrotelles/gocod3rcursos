package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço da variável e atribuindo
	*p++   // desreferenciando (pegando o valor)
	i++
	fmt.Println(p, *p, i, &i)
	
	// Go não tem aritmética de ponteiros
	// p++
}
