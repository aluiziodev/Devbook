package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

func VerificarSenha(senhaWhash, senhaString string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(senhaWhash), []byte(senhaString)); err != nil {
		return errors.New("Senha Incorreta !! ")
	}
	return nil
}
