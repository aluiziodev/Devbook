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
		Uri:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarPublicacao,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarPublicacao,
		RequerAuth: true,
	},
	{
		Uri:        "/usuarios/{usuarioId}/publicacoes",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacoesUsuario,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoes/{publicacaoId}/curtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CurtirPublicacao,
		RequerAuth: true,
	},
	{
		Uri:        "/publicacoes/{publicacaoId}/curtir",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DescurtirPublicacao,
		RequerAuth: true,
	},
}
