package main

import (
	"log"
)

var port = ":3000"

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(port, store)
	server.Run()
}
