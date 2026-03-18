package controllers

import (
	"apiDevbook/src/database"
	"apiDevbook/src/models"
	"apiDevbook/src/models/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario insere no DB
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario models.Usuario
	if err = json.Unmarshal(bodyReq, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := database.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	repositorio := repositorios.NovoRepoUsuarios(db)
	repositorio.Criar(usuario)

}

// BuscarUsuarios faz uma query por todos os usurios no DB
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

}

// BuscarUsuario faz uma query por um usuario especifico no DB
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}

// AtualizarUsuario faz um update em um usuario especifico no DB
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

}

// DeletarUsuario faz o delete de um usuario especifico no DB
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

}
