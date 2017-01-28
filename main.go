package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andrewmelis/dadzip/ziphandler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/filenames", ziphandler.MultiPartZipHandler)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	log.Printf("starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
