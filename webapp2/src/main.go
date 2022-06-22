package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8443",
		Handler: nil,
	}
	err := server.ListenAndServeTLS("../cert/cert.pem", "../cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
