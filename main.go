package main

import (
	"learn-rest-api/cmd/app/route"
	"learn-rest-api/config"
)

func main() {
	config.Init()
	config.InitDB()

	var router = route.Router()
	router.Run()
}