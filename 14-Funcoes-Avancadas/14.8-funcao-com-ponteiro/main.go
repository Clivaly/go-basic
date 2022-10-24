package main

import "fmt"

// Sem ponteiro
func inverterSinal(numero int) int {
	return numero * -1
}

// Com ponteiro
func inverterSinalComPonteiro(numero *int) { // Não precisa de retorno, pois está alterando direto na variável
	// *numero = *numero * -1
	*numero *= -1 // forma simplificada
}

func main() {
	num := 20
	invertido := inverterSinal(num)
	fmt.Println(invertido)
	fmt.Println(num)

	num2 := 40
	fmt.Println(num2)
	inverterSinalComPonteiro(&num2)
	fmt.Println(num2)

}
