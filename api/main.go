package main

import (
	"apiDevbook/src/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Iniciando a API!!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8000", r))

}
