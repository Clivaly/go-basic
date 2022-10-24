package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("for")
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i:", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println("Chegou a:", i)
	fmt.Println("------------------------")

	// Outro exemplo For com init, porem, assim a variavel fica apenas dentro local do For
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("subindo numero de j:", j)
	// 	time.Sleep(time.Second)
	// }

	// For com range
	nomes := [3]string{"Pedro", "Jose", "Lucas"}
	
	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}
	
	// For com range, outro exemplo
	for _, valor := range "PALAVRA" {
		fmt.Println(string(valor))
	}
	
	// For com range, outro exemplo, usando maps
	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Souza",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}
	
	// For infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)		
	}
}
