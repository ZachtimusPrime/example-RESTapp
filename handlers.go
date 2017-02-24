package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"
const notImplementedError = "Method is not implemented on this endpoint"

type message struct {
	Msg string `json:"message"`
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

	hello := &message{Msg: "hello!"}

	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(hello)
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(501)
	fmt.Fprint(w, notImplementedError)
}
