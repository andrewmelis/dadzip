package main

import (
	"log"
	"net/http"

	"github.com/andrewmelis/dadzip/ziphandler"
)

func main() {
	log.Printf("starting server...")
	http.HandleFunc("/filenames", ziphandler.ZipHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
