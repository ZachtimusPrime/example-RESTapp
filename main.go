package main

import (
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Setup router mux
	router := mux.NewRouter().StrictSlash(false)

	// Define routes, methods, and handlers
	router.HandleFunc("/hello", sayHelloHandler).Methods("GET")
	router.HandleFunc("/hello", notImplementedHandler).Methods("POST","PUT","PATCH","DELETE","HEAD","OPTIONS")

	// Start server
	ServicePort := "8080"
	portConfig := os.Getenv("HTTP_PLATFORM_PORT")
	if portConfig != "" {
		ServicePort = portConfig
	}
	log.Println("Starting server on port " + ServicePort)
	http.ListenAndServe(":"+ServicePort,router)
}
