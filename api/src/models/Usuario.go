package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id         uint64    `json:"id,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	DataInicio time.Time `json:"dataInicio,omitempty"`
}

func (user *Usuario) Validar() error {
	if user.Nome == "" {
		return errors.New("O nome é obrigatório, não pode estar em branco!! ")
	}

	if user.Email == "" {
		return errors.New("O email é obrigatório, não pode estar em branco!! ")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatório, não pode estar em branco!! ")
	}

	if user.Password == "" {
		return errors.New("A senha é obrigatório, não pode estar em branco!! ")

	}

	return nil
}

func (user *Usuario) Formatar() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

}

// Valida e Formata usuario
func (user *Usuario) Preparar() error {
	if err := user.Validar(); err != nil {
		return errors.New(err.Error())
	}
	user.Formatar()
	return nil
}
