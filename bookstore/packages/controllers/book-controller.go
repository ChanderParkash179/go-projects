package controllers

import (
	"bookstore/packages/models"
	"bookstore/packages/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var bookRequest models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := models.GetBooks()

	err := utils.APIResponse(w, http.StatusOK, books)
	if err != nil {
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return
	}

	fmt.Printf("ID: %+v\n", id)

	book, _ := models.GetBook(id)

	fmt.Printf("Book: %+v\n", &book)

	err = utils.APIResponse(w, http.StatusOK, book)
	if err != nil {
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return
	}

	book, _ := models.DeleteBook(id)

	err = utils.APIResponse(w, http.StatusOK, book)
	if err != nil {
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}

	err := utils.Decode(r, CreateBook)
	if err != nil {
		return
	}

	body, _ := CreateBook.CreateBook()

	err = utils.APIResponse(w, http.StatusOK, body)
	if err != nil {
		return
	}
}
