package routes

import (
	"github.com/MRafaydev/go-bookstore/packages/contollers"
	"github.com/gorilla/mux"
)

var RegisterBooksStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", contollers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", contollers.GetBook).Methods("GET")
	router.HandleFunc("/books/{booksid}", contollers.GetBooksByID).Methods("GET")
	router.HandleFunc("/books/{booksid}", contollers.UpdateBooks).Methods("PUT")
	router.HandleFunc("/books/{booksid}", contollers.DeleteBooks).Methods("DELETE")
}
