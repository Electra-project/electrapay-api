package main

import (
	"github.com/joho/godotenv"
	"github.com/ruannelloyd/electrapay-api/src/helpers"
	"github.com/ruannelloyd/electrapay-api/src/server"
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
