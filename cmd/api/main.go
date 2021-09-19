package main

import (
	"fmt"
	"log"

	"github.com/ZombieMInd/slovodel/internal/slovodel/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := &server.Config{}

	server.InitConfig(conf)

	fmt.Printf("Starting %s", conf.Name)
	server.Start(conf)
}
