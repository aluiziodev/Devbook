package controllers

import (
	"apiDevbook/src/database"
	"apiDevbook/src/models"
	"apiDevbook/src/models/repositorios"
	"apiDevbook/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CriarUsuario insere no DB
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(bodyReq, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repositorio := repositorios.NovoRepoUsuarios(db)
	usuario.Id, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios faz uma query por todos os usurios no DB
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nickOrName := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	usuarios, err := repositorio.Buscar(nickOrName)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

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
