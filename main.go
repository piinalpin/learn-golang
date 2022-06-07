package main

import (
	"learn-rest-api/cmd/app/route"
	"learn-rest-api/config"
)

func main() {
	init := config.Init()
	var router = route.Router(init)
	router.Run()
}