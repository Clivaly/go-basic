package main

import "fmt"

func main() {
	func() {
		fmt.Println("Ola Mundo!")
	}() //executa imediatamente a função anônima

	// Exemplo com parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmentro") //executa imediatamente a função anônima

	// Exemplo com retorno e armazenando em uma variável
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmentro 2") //executa imediatamente a função anônima

	fmt.Println(retorno)
}

