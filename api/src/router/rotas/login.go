package rotas

import (
	"apiDevbook/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Uri:        "/login",
	Metodo:     http.MethodPost,
	Funcao:     controllers.Login,
	RequerAuth: false,
}
