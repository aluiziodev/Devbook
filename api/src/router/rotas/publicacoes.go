package rotas

import (
	"apiDevbook/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		Uri:        "/publicacoes",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarPublicacao,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoes",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacoes,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacao,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoe/{publicacaoId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarPublicacao,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoes",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarPublicacao,
		RequerAuth: true,
	},
}
