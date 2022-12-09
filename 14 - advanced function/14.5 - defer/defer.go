package main

import "fmt"

func funcao1() {
	fmt.Println("funcao 1")
}

func funcao2() {
	fmt.Println("funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("media calculada, resultado sera retornada")
	fmt.Println("start aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}
