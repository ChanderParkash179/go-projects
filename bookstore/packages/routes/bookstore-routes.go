package routes

import (
	"bookstore/packages/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

func (s *APIServer) Run() {

	r := mux.NewRouter()

	r.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	//r.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")

	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")

	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")

	log.Printf("\nserver is running on port: %s\n", s.listenAddress)

	err := http.ListenAndServe(s.listenAddress, r)
	if err != nil {
		return
	}
}
