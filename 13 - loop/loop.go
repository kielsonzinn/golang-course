package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("for")

	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i", i)
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 5 {
		fmt.Println("incrementando j", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"Joao", "Davi", "Luiz"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "WORD" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("infinite")
		time.Sleep(time.Second)
	}

}
