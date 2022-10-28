package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Sincronizar Goroutines
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Especificar a quantidade de goroutines

	go func()  {
		escrever("Olá Mundo!") 
		waitGroup.Done() // -1 para esperar
	}()

	go func()  {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Esperar a contagem das goroutines chegar em zero.

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}