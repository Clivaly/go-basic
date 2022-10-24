package main

import (
	"fmt"
	"math"
)

type forma interface { // Interface só tem assinaturas de métodos
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

// Implementando o método "area()" da interface "forma"
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	// return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2) // Outra maneira
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
