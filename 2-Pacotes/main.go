package main

import (
	"fmt"

	"github.com/go-basic/2-Pacotes/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.") // Portuguese
	fmt.Println("Writing from main file.")     // English
	auxiliar.Escrever()
	// auxiliar.escrever()
	// Erro pois a função escrever() não está exportada/pública.
	// Error because the escrever() function is not exported/public.
}
