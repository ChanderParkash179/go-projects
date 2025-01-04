package main

import (
	"bookstore/packages/models"
	"bookstore/packages/routes"
)

var port = ":4000"

func main() {

	models.Init()

	server := routes.NewAPIServer(port)
	server.Run()
}
