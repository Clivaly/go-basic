package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	// mensagem := <- canal
	// fmt.Println(mensagem)

	// Exemplo com Loop infinito
	// for { // Tome cuidado, assim causa deadlock, que só é pego este erro em execução e não em compilação!
	// 	mensagem, aberto := <- canal
	// 	if !aberto { // Verificar se o canal está aberto
	// 		break // Sair do loop infinito
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// Alternativa
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) // Para fechar o canal
}
