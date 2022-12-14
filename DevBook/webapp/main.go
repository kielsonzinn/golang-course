package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

/**
func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)
	fmt.Println(blockKey)
}
**/

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Rodando WebApp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
