package rotas

import (
	"apiDevbook/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		Uri:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		RequerAuth: false,
	},
	{
		Uri:        "/usuarios/{id}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: false,
	},
	{
		Uri:        "/usuarios/{id}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: false,
	},
	{
		Uri:        "/usuarios/{id}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: false,
	},
}
