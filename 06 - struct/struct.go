package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	fmt.Println("struct")

	var usuario1 usuario
	usuario1.nome = "Davi"
	usuario1.idade = 21

	fmt.Println(usuario1)

	endereco1 := endereco{"Rua dos bobo", 0}

	usuario2 := usuario{"Davi", 21, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)

}
