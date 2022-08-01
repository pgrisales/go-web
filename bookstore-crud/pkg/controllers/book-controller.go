package controllers

import (
	"net/http"

	"github.com/pgrisales/go-web/bookstore-crud/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
}
