package controllers

import (
	"apiDevbook/src/authentication"
	"apiDevbook/src/database"
	"apiDevbook/src/models"
	"apiDevbook/src/repositorios"
	"apiDevbook/src/respostas"
	"apiDevbook/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepoUsuarios(db)
	usuarioBanco, err := repositorio.BuscarEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuarioBanco.Password, usuario.Password); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CriarToken(usuarioBanco.Id)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.Menssagem(w, http.StatusOK, fmt.Sprintf("Usuario logado com sucesso!! token: %s", token))

}
