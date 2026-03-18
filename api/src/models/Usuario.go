package models

import "time"

type Usuario struct {
	Id         uint64    `json:"id,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	DataInicio time.Time `json:"dataInicio,omitempty"`
}
