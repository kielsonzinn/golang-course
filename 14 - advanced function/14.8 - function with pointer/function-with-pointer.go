package main

import "fmt"

func inverterSignal(numero int) int {
	return numero * -1
}

func inverterSignalWithPointer(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSignal(numero)

	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40

	inverterSignalWithPointer(&novoNumero)
	fmt.Println(novoNumero)
}
