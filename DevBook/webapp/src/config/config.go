package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//ApiUrl representa a URL para comunicação com a API
	APIURL = ""

	//Porta onde a aplicação web está rodando
	Porta = 0

	//HashKey é utilizada para autenticar o cookie
	HashKey []byte

	//BlockKey é utilizda para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		Porta = 3000
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
