package routes

import (
	"github.com/gorilla/mux"
	"github.com/nesuh/Book_Managmen/pkg/controllers"
)

var ReisterBooksRoutes = func(router *mux.Router) {
	router.HandleFunc("book/", controllers.CreateBook()).Methods("POST")
	router.HandleFunc("book/", controllers.GetBook()).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.GetBookById()).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.UpdateBook()).Methods("PUT")
	router.HandleFunc("book/{bookId}", controllers.DeleteBook()).Methods("DELETE")
}
