package controllers

import (
	"apiDevbook/src/authentication"
	"apiDevbook/src/database"
	"apiDevbook/src/models"
	"apiDevbook/src/repositorios"
	"apiDevbook/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = usuario.Preparar("cadastro"); err != nil {
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
	parametros := mux.Vars(r)

	userId, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	usuario, err := repositorio.BuscarId(userId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// AtualizarUsuario faz um update em um usuario especifico no DB
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIdToken, err := authentication.ExtrairUserId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userIdToken != id {
		respostas.Erro(w, http.StatusForbidden, errors.New("Nao 'e possivel atualizar este usuario!!"))
		return
	}

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

	if err = usuario.Preparar("atualizar"); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	if err = repositorio.Atualizar(id, usuario); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

// DeletarUsuario faz o delete de um usuario especifico no DB
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIdToken, err := authentication.ExtrairUserId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userIdToken != id {
		respostas.Erro(w, http.StatusForbidden, errors.New("Nao 'e possivel deletar este usuario!!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	if err = repositorio.Deletar(id); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	userIdSeguidor, err := authentication.ExtrairUserId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	userIdSeguir, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userIdSeguidor == userIdSeguir {
		respostas.Erro(w, http.StatusForbidden, errors.New("Nao 'e possivel seguir a si mesmo!!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	if err = repositorio.Seguir(userIdSeguir, userIdSeguidor); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func PararSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	userIdSeguidor, err := authentication.ExtrairUserId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	userIdSeguir, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userIdSeguidor == userIdSeguir {
		respostas.Erro(w, http.StatusForbidden, errors.New("Nao 'e possivel deixar de seguir a si mesmo!!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	if err = repositorio.DeixarSeguir(userIdSeguir, userIdSeguidor); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	seguidores, err := repositorio.BuscarSeguidores(userId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)

}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	seguindo, err := repositorio.BuscarSeguindo(userId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, seguindo)

}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	userIdToken, err := authentication.ExtrairUserId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	userId, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userIdToken != userId {
		respostas.Erro(w, http.StatusForbidden, errors.New("Nao 'e possivel atualizar a senha desse usuario!!"))
		return
	}

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var senha models.Password
	if err = json.Unmarshal(bodyReq, &senha); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

}
