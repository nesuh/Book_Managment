package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nesuh/Book_Managment/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.ReisterBooksRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
