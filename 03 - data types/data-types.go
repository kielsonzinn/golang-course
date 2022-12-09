package main

import (
	"errors"
	"fmt"
)

func main() {

	//NUMEROS INTEIROS

	var numInt int8 = 10
	fmt.Println(numInt)

	var numUint uint32 = 10000
	fmt.Println(numUint)

	// alias
	// INT32 = RUNE
	var numRune rune = 12456
	fmt.Println(numRune)

	// BYTE = UINT8
	var numByte byte = 123
	fmt.Println(numByte)

	//FIM NUMEROS INTEIROS

	//NUMEROS REAIS

	var numFloat float32 = 123.45
	fmt.Println(numFloat)

	var numFloat64 float64 = 12356.45
	fmt.Println(numFloat64)

	// FIM NUMERO REAIS

	// STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// FIM STRINGS

	//NAO INICIALIAR

	var texto int16
	fmt.Println(texto)

	//FIM NAO INICIALIZAR

	//BOOLEAN

	var bool1 bool
	fmt.Println(bool1)

	//FIM BOOLEAN

	//ERROR

	var erro error = errors.New("internal error")
	fmt.Println(erro)

	//FIM ERROR

}
