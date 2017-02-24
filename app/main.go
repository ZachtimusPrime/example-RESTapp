package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	http.HandleFunc("/", http.FileServer(http.Dir("./apidocs/")))
}

func main() {
	// Setup router mux
	router := mux.NewRouter().StrictSlash(false)

	// Define routes, methods, and handlers
	router.Path("/").Handler(http.FileServer(http.Dir("./apidocs/")))

	router.HandleFunc("/hello", sayHelloHandler).Methods("GET")
	router.HandleFunc("/hello", notImplementedHandler).Methods("POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS")
	
}
