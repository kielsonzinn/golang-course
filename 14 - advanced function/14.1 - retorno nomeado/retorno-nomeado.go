package main

import "fmt"

func calcMath(n1, n2 int) (soma int, subtracao int) {

	soma = n1 + n2
	subtracao = n1 - n2

	return

}

func main() {
	fmt.Println("retorno nomeado")
	fmt.Println(calcMath(1, 3))
}
