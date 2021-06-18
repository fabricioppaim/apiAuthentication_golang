package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriaTokenv retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{}
	permissoes["autorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.Secretkey))
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")
}

// extrairToken verifica se o token está no formato válido e extrai ele do header
func extrairToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// retornarChaveDeVerificacao verifica se a chave foi gerado pelo mesmo método que o token foi criado.
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

		return nil, fmt.Errorf("método de assinatura inesperado! %w", token.Header["alg"])

	}

	return config.Secretkey, nil
}