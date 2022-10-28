package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     uint   `json:"idade"`
}

func main() {
	// Exemplo 1, converte JSON em struct
	pessoaEmJSON := `{"nome":"Joao","sobrenome":"Silva","idade":32}`
	// Passar os valores para a variável "p"
	var p pessoa
	if erro := json.Unmarshal([]byte(pessoaEmJSON), &p); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(p)

	// Exemplo 2, converte JSON em map
	pessoa2EmJSON := `{"nome":"Joao","sobrenome":"Silva"}`
	// Passar os valores para a variável "p2"
	var p2 = make(map[string]string)
	if erro := json.Unmarshal([]byte(pessoa2EmJSON), &p2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(p2)

}