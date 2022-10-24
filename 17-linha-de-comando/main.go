package main

import (
	"log"
	"os"

	"linha-de-comando/app"
)

func main() {
	aplicacao := app.Gerar()
	err := aplicacao.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
