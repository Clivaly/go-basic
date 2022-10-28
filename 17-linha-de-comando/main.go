package main

import (
	"log"
	"os"

	"github.com/go-basic/17-linha-de-comando/app"
)

func main() {
	aplicacao := app.Gerar()
	err := aplicacao.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
