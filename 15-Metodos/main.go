package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// Método 1
func (u usuario) salvar() {
	fmt.Printf("Salvando usuário: %s, com idade: %d, no banco de dados\n", u.nome, u.idade)
}

// Método 2
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Método 3 com ponteiro
func (u *usuario) fazerAniversario() { // sem o ponteiro não irá alterar a idade
	u.idade++
}

func main() {
	user1 := usuario{nome: "Pedro", idade: 5}
	user1.salvar()
	m := user1.maiorDeIdade()
	fmt.Println(m)
	user1.fazerAniversario()
	fmt.Println(user1.idade)
}
