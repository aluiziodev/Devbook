package authentication

import (
	"apiDevbook/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CriarToken(userId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString(config.SecretKey)
}

func ValidarToken(r *http.Request) error {
	tokenStr := extrairToken(r)
	if tokenStr == "" {
		return errors.New("Erro ao validar token")
	}

	token, err := jwt.Parse(tokenStr, returnKeyDeVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido!!")

}

func ExtrairUserId(r *http.Request) (uint64, error) {
	tokenStr := extrairToken(r)
	if tokenStr == "" {
		return 0, errors.New("Erro ao validar token")
	}

	token, err := jwt.Parse(tokenStr, returnKeyDeVerificacao)
	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return 0, errors.New("Token invalido!!")

}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKeyDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New(fmt.Sprintf("Metodo de assinatura inesperado!! %v",
			token.Header["alg"]))
	}

	return config.SecretKey, nil
}
