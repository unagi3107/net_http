package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	err := server.ListenAndServeTLS("webapp2/cert/cert.pem", "webapp2/cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
