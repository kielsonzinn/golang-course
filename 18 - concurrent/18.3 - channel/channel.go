package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)

	go escrever("hello world", canal)
	fmt.Println("depois da função escrever ser chamada")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("end of program")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- fmt.Sprintf("%s %d", texto, i)
		time.Sleep(time.Second)
	}

	close(canal)
}
