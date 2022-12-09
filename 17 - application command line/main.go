package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("start point")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
