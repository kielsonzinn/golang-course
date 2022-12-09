package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calcs(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {

	soma := somar(1, 5)
	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println(text)
		return text
	}

	text := f("call function f")
	fmt.Println(text)

	_, resultadoSubtracao := calcs(10, 15)
	fmt.Println(resultadoSubtracao)

}
