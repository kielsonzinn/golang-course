package main

import (
	"fmt"
)

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s %d", texto, 10)
	}("hello world!")

	fmt.Println(retorno)

}
