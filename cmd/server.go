package main

import (
	"log"
	"net/http"
	"github.com/anton-bryukhov/url-shortener/pkg/handler"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
