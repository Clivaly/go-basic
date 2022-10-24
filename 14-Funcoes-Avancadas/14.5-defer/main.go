package main

import "fmt"

func f1()  {
	fmt.Println("Executando f1")	
}

func f2()  {
	fmt.Println("Executando f2")	
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}

	return false
}

func main() {	
	// DEFER = ADIAR 
	defer f1()
	f2()

	fmt.Println(alunoAprovado(7, 8))
}