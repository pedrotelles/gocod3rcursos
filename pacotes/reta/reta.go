package main

import "math"

// Iniciando com letra musica é PÚBLICO (vísivel dentro e fora de um pacote)!

// Inicializando com letra minúscula é PRIVADO ( vísivel dentro do pacote)!

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distância linear entre os dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
