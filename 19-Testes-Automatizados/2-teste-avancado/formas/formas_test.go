package formas_test

import (
	"math"
	"testing"

	. "github.com/go-basic/19-Testes-Automatizados/2-teste-avancado/formas"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areRecebida := ret.Area()

		if areaEsperada != areRecebida {
			t.Fatalf("A area recebida: %f, é diferente da area esperada: %f",
				areRecebida, areaEsperada,
			)
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{Raio: 10}
		raioEsperado := float64(math.Pi * 100)
		raioRecebido := circ.Area()

		if raioEsperado != raioRecebido {
			t.Fatalf("O raio recebido: %f, é diferente do raio esperado: %f",
				raioRecebido, raioEsperado,
			)
		}

	})
}
