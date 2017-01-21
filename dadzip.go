package dadzip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func zipHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	// validate that sent file is zip

	zr, err := zip.NewReader(bytes.NewReader(body), r.ContentLength)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range zr.File {
		fmt.Fprintf(w, "%s\n", f.FileHeader.Name)
	}
	// fmt.Fprintf(w, "foobar\n")
}
