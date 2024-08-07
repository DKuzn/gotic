package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))

	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", handlers.CORS(origins, methods)(r))
}
