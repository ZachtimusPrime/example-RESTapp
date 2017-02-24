package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

const content_type = "Content-Type"
const application_json = "application/json"
const notImplementedError = "Method is not implemented on this endpoint"

type Message struct {
	Msg string `json:"message"`
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

	hello := &Message{Msg: "hello!"}

	// return the json list of movie names
	w.Header().Set(content_type, application_json)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(hello)
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content_type, application_json)
	w.WriteHeader(501)
	fmt.Fprint(w, notImplementedError)
}