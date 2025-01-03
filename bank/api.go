package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// server setup
type APIServer struct {
	listenAddress string
	store         Storage
}

func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

// Run
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.handleAccounts))
	router.HandleFunc("/accounts/del", makeHTTPHandleFunc(s.handleDeleteAccounts))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleCreateAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Printf("\nserver is running on port: %s\n", s.listenAddress)

	err := http.ListenAndServe(s.listenAddress, router)
	if err != nil {
		return
	}
}

// handlers
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Chander", "Parkash")
	return APIResponse(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	request := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}

	acc := NewAccount(request.FirstName, request.LastName)
	if err := s.store.CreateAccount(acc); err != nil {
		return err
	}

	return APIResponse(w, http.StatusCreated, acc)
}

func (s *APIServer) handleAccounts(w http.ResponseWriter, r *http.Request) error {
	resp, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return APIResponse(w, http.StatusOK, resp)
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccounts(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// API Generic Handler
type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func APIResponse(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			APIResponse(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
