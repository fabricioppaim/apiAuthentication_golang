package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// StringConexaoBanco é a string de conexão com o MySQL
	StringConexaoBanco = ""

	// Porta é onde a API vai expor a conexão
	Porta = 0

	// Secretkey é a chave usada para assinar o token
	Secretkey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAMEDB"),
		os.Getenv("DB_LOCALZONE"),
	)

	Secretkey = []byte(os.Getenv("SECRET_KEY"))
}
