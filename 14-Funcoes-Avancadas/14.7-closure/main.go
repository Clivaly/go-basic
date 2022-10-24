package main

import "fmt"

// Closure: são funções que referenciam variáveis fora do seu corpo.
func closure() func() { // retornando uma função
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova() 
}