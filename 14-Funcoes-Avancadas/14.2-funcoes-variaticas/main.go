package main

import "fmt"

// Varios numeros podem ser passados
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Outro exemplo
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	totalDaSoma := soma(1, 2, 3)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 1, 2, 3, 4, 5)
}
