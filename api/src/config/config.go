package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConnDB = ""
	Port         = "0"
)

func Carregar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	Port = os.Getenv("PORT")

	StringConnDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}
