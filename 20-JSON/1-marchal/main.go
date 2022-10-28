package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Exemplo 1, converte struct em JSON
	c := cachorro{"Rex", "Buldog", 3}
	// Converter struct para JSON
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)
	// Converter a saida
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	// Exemplo 2, converte maps em JSON
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	// Converter a saida
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
