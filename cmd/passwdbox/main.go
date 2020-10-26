package main

import (
	"flag"
	"log"

	"github.com/passwdapp/box/config"
)

func main() {
	useDotenv := flag.Bool("use-dotenv", true, "Use dotenv instead of environment variables")
	flag.Parse()

	log.Println("Initializing passwd server")
	config.InitConfig(*useDotenv)
}
