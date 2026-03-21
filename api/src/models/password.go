package models

type Password struct {
	Nova  string `json:"new"`
	Atual string `json:"old"`
}
