package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("./apidocs/")))
}

func main() {
	// Setup router mux
	router := mux.NewRouter().StrictSlash(false)

	// Define routes, methods, and handlers
	router.Path("/").Handler(http.FileServer(http.Dir("./apidocs/")))

	router.HandleFunc("/hello", sayHelloHandler).Methods("GET")
	router.HandleFunc("/hello", notImplementedHandler).Methods("POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS")

}
