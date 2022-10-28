package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrência != paralelismo
	go escrever("Olá Mundo!") // goroutine, executa a função e passa para a proxima...
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}