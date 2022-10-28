package formas

import (
	"math"
)

type Forma interface { // Interface só tem assinaturas de métodos
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

// Implementando o método "area()" da interface "forma"
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	// return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.Raio, 2) // Outra maneira
}

// func EscreverArea(f Forma) {
// 	fmt.Printf("A area da forma é %0.2f\n", f.Area())
// }
