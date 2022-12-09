package main

import "fmt"

func main() {

	fmt.Println("if else")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")

	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")

	} else if outroNumero < -10 {
		fmt.Println("numero é menor quue -10")

	} else {
		fmt.Println("entre 0 e -10")
	}

}
