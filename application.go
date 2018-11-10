package main

import (
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/server"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	helpers.LogInfo("Starting server...")
	server.Start()
}
