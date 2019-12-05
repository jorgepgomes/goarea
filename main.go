package goarea

import "math"

// Pi é uma proporção númerica definida pela relação entre
// o perimetro de uma circuferencia e seu diametro
const Pi = 3.1416

// Circ é responsável por caular a área da circuferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
