package models

import (
	"apiDevbook/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	Id         uint64    `json:"id,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	DataInicio time.Time `json:"dataInicio,omitempty"`
}

func (user *Usuario) Validar(etapa string) error {
	if user.Nome == "" {
		return errors.New("O nome é obrigatório, não pode estar em branco!! ")
	}

	if user.Email == "" {
		return errors.New("O email é obrigatório, não pode estar em branco!! ")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email inserido é invalido!!")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatório, não pode estar em branco!! ")
	}

	if etapa == "cadastro" && user.Password == "" {
		return errors.New("A senha é obrigatório, não pode estar em branco!! ")

	}

	return nil
}

func (user *Usuario) Formatar(etapa string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		senhaWhash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(senhaWhash)
	}

	return nil

}

// Valida e Formata usuario
func (user *Usuario) Preparar(etapa string) error {
	if err := user.Validar(etapa); err != nil {
		return err
	}
	if err := user.Formatar(etapa); err != nil {
		return err
	}
	return nil
}
