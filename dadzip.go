package dadzip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"path/filepath"
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
		info := f.FileHeader.FileInfo()
		if !info.IsDir() {
			fmt.Fprintf(w, "%s\n", filepath.Base(info.Name()))
			continue
		}

	}
}
