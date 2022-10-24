package main

import "fmt"

func main() {
	fmt.Println("Maps")
	
	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	
	fmt.Println(usuario["nome"])
	
	// Map aninhado
	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": "Silva",
		},
		"curso": {
			"escolha-1": "Matematica",
			"escolha-2": "Portugues",
			"escolha-3": "Quimica",
		},
	}
	
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	
	usuario2["signo"] = map[string]string {
		"nome": "libra",
	}
	fmt.Println(usuario2)
}