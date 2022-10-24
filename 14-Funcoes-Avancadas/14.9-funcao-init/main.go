package main

import "fmt"

// Exemplo, inicialiar uma variavel
var n int

func init() { // sem a função init, a variável "n" não seria executada pela função main
	fmt.Println("Função init é executada antes da main")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
