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
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}/parar-seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.PararSeguirUsuario,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}/seguidores",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}/seguindo",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguindo,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{id}/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.AtualizarSenha,
		RequerAuth: true,
	},
}
