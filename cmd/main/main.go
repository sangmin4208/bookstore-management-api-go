package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sangmin4208/bookstore-management-api-go/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
