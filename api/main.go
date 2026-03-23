package main

import (
	"apiDevbook/src/config"
	"apiDevbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()

	r := router.Gerar()
	log.Printf("Iniciando a API na porta: %s", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r))

}
