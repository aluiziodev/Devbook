package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	Id          uint64    `json:"id,omitempty"`
	Titulo      string    `json:"titulo,omitempty"`
	Conteudo    string    `json:"conteudo,omitempty"`
	AutorId     uint64    `json:"autorId,omitempty"`
	AutorNick   string    `json:"autorNick,omitempty"`
	Curtidas    uint64    `json:"curtidas"`
	DataCriacao time.Time `json:"dataCriacao,omitempty"`
}

func (publi *Publicacao) Preparar() error {
	if err := publi.validar(); err != nil {
		return err
	}
	publi.formatar()
	return nil
}

func (publi *Publicacao) validar() error {
	if publi.Titulo == "" {
		return errors.New("O titulo é um campo obrigatório, não pode estar vazio!!")
	}
	if publi.Conteudo == "" {
		return errors.New("O conteudo é um campo obrigatório, não pode estar vazio!!")
	}

	return nil
}

func (publi *Publicacao) formatar() {
	publi.Titulo = strings.TrimSpace(publi.Titulo)
	publi.Conteudo = strings.TrimSpace(publi.Conteudo)
}
