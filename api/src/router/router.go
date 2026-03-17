package router

import (
	"apiDevbook/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	router := mux.NewRouter()
	rotas.Configurar(router)
	return router

}
