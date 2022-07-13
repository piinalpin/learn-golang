package main

import (
	"learn-rest-api/cmd/app/route"
	"learn-rest-api/config"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	init := config.Init()
	var router = route.Router(init)
	router.Run(":" + port)
}
