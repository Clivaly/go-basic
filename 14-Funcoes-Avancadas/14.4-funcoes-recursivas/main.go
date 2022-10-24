package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main() {	
	fmt.Println("Funções Recursivas")

	p1 := uint(5)
	fmt.Println(fibonacci(p1))
	fmt.Println("-----------")

	// Exemplo com for
	p2 := uint(7)

	for i := uint(0); i < p2; i++ {
		fmt.Println(fibonacci(i))
	}
}