package main

import (
	"fmt"

	"github.com/go-basic/19-Testes-Automatizados/1-introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
