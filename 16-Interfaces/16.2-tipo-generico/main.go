package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(1)
	generica(1.2)
	generica(true)
	generica("Pedro")
	fmt.Println("---------------------")

	// Caso não recomendado, um "map" recebendo interface generica, exemplo:
	mapa := map[interface{}]interface{} {
		1: "string",
		float64(100): false,
		true: float32(50),
		"string": 17.32,
	}

	fmt.Println(mapa)
}
