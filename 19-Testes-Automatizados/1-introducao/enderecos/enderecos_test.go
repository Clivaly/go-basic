// TESTE DE UNIDADE

package enderecos_test

import (
	"testing"

	// Com este ".", não precisa colocar o pacote para acessar a função
	. "github.com/go-basic/19-Testes-Automatizados/1-introducao/enderecos"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipodeEndereco(t *testing.T) {
	// Essa função roda os testes em paralelo em vez de um de cada vez, colocar em todas as funções
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{enderecoInserido: "Rua go", retornoEsperado: "Rua"},
		{enderecoInserido: "Avenida Paulista", retornoEsperado: "Avenida"},
		{enderecoInserido: "Estrada 123", retornoEsperado: "Estrada"},
		{enderecoInserido: "Rodovia dos Imigrantes", retornoEsperado: "Rodovia"},
		{enderecoInserido: "", retornoEsperado: "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido: %s, é diferente do esperado: %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
