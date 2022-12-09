package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println(config.APIURL)
	fmt.Println(config.Porta)
	fmt.Println(config.HashKey)
	fmt.Println(config.BlockKey)

	fmt.Printf("Rodando WebApp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
