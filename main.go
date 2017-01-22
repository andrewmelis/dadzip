package main

import (
	"log"
	"net/http"

	"github.com/andrewmelis/dadzip/ziphandler"
)

func main() {
	log.Printf("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/filenames", ziphandler.MultiPartZipHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
