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
	SecretKey    = []byte{}
)

func Carregar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	Port = os.Getenv("PORT")
	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	StringConnDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}
