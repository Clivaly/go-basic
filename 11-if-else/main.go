package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 0

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Igual a 15")		
	} else {
		fmt.Println("Menor que 15")
	}

	// No if init, a variavel tem scopo somente local dentro do if/else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero!")
	} else {
		fmt.Println("Numero é menor que zero!")
	}
}
