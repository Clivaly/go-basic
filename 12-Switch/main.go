package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "Numero inválido!"
	}
}

// Outro exemplo 
func diaSemana2(numero int) string {
	var diaDaSemana string

	switch  {
	case numero == 1:
		diaDaSemana = "domingo"
	case numero == 2:
		diaDaSemana = "segunda"
	case numero == 3:
		diaDaSemana = "terça"
	case numero == 4:
		diaDaSemana = "quarta"
	case numero == 5:
		diaDaSemana = "quinta"
	case numero == 6:
		diaDaSemana = "sexta"
	case numero == 7:
		diaDaSemana = "sabado"
	default:
		diaDaSemana = "Numero inválido!"
	}

	return diaDaSemana
}
func main() {
	fmt.Println("Switch")

	dia := diaSemana(1)
	fmt.Println(dia)


}
