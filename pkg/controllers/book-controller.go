package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sangmin4208/bookstore-management-api-go/pkg/models"
	"github.com/sangmin4208/bookstore-management-api-go/pkg/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		fmt.Println(err)
	}
	data, err := json.Marshal(books)
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Fprintf(w, "Invalid book id")
		return
	}
	book, err := models.GetBookById(uint(id))
	if err != nil {
		fmt.Fprintf(w, "The book with given id was not found")
		return
	}
	data, err := json.Marshal(book)
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	b, err := book.CreateBook()
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	data, err := json.Marshal(b)
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Fprintf(w, "Invalid book id")
		return
	}
	book, err := models.DeleteBook(uint(id))
	if err != nil {
		fmt.Fprintf(w, "The book with given id was not found")
		return
	}
	data, err := json.Marshal(book)
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Fprintf(w, "Invalid book id")
		return
	}
	var updateBook models.Book
	utils.ParseBody(r, &updateBook)
	updateBook.ID = uint(id)
	bookDetail, err := models.GetBookById(uint(id))
	if err != nil {
		fmt.Fprintf(w, "The book with given id was not found")
		return
	}
	if updateBook.Name != "" {
		updateBook.Name = bookDetail.Name
	}
	if updateBook.Author != "" {
		updateBook.Author = bookDetail.Author
	}
	if updateBook.Publication != "" {
		updateBook.Publication = bookDetail.Publication
	}
	b, err := updateBook.UpdateBook()
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	data, err := json.Marshal(b)
	if err != nil {
		fmt.Fprintf(w, "An error occurred while processing the request")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}
