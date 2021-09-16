package main

import (
	"fmt"
	"log"

	"github.com/ZombieMInd/slovodel/internal/slovodel/config"
	"github.com/ZombieMInd/slovodel/internal/slovodel/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &config.App{}

	server.InitConfig(app)

	fmt.Printf("Starting %s", app.Name)
}
