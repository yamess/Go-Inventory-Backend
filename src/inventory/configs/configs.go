package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Host string
var PgDbUrl string
var Version string
var DefaultUser uint = 1

func InitEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Error while loading the environment file")
	}
	Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	PgDbUrl = os.Getenv("DB_URL")
	Version = os.Getenv("VERSION")
}
