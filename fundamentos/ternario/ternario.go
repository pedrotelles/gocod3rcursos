package main

import "fmt"

// Não tem operador ternário
func ObterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(ObterResultado(6.2))
}
