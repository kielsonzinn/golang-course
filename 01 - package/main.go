package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("write from main file")
	auxiliar.Write()
	erro := checkmail.ValidateFormat("golang@gmail.com")
	fmt.Println(erro)
}
