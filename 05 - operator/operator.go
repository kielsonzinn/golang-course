package main

import "fmt"

func main() {

	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 4

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//FIM ARITMETICOS

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LOGICOS
	fmt.Println("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM DOS OPERADORES LOGICOS

	//OPERADORES UNARIOS

	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 //numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)
	//FIM DOS OPERADORES UNARIOS

	//OPERADOR TERNARIO
	var texto string
	if numero > 1 {
		texto = "maior que 1"

	} else if numero == 1 {
		texto = "igual a 1"

	} else {
		texto = "menor que 1"
	}

	fmt.Println(texto)

}
