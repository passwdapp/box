package main

import (
	"flag"
	"log"

	"github.com/passwdapp/box/config"
	"github.com/passwdapp/box/database"
	"github.com/passwdapp/box/http"
)

func main() {
	useDotenv := flag.Bool("use-dotenv", true, "Use dotenv instead of environment variables")
	flag.Parse()

	cfg := config.Config{}

	log.Println("Initializing passwd server")
	cfg.InitConfig(*useDotenv)

	log.Println("Connecting to the database")
	database.Connect(cfg.DatabaseFile)

	log.Println("Initializing the HTTP server")
	http.InitHTTP(&cfg)
}
