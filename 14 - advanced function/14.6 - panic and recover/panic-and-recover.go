package main

import "fmt"

func recuperarExecucacao() {
	if r := recover(); r != nil {
		fmt.Println("recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {

	defer recuperarExecucacao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}

	if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6!")

}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("pos execution")
}
