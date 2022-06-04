package main

import (
	"learn-rest-api/cmd/app/route"
)

func main() {
	var router = route.Router()
	router.Run()
}