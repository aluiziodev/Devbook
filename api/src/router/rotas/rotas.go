package rotas

import (
	"apiDevbook/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri        string
	Metodo     string
	Funcao     func(http.ResponseWriter, *http.Request)
	RequerAuth bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		if rota.RequerAuth {
			r.HandleFunc(rota.Uri,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
