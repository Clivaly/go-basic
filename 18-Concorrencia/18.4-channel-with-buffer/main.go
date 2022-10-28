package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Capacidade | Buffer
	canal <- "Olá Mundo!"
	canal <- "Programando em GO!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}