package ziphandler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func ZipHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received: %+v", r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s: no body found\n", err)
		fmt.Fprintf(w, "%s: no body found\n", err)
		return
	}
	defer r.Body.Close()

	// validate that sent file is zip
	zr, err := zip.NewReader(bytes.NewReader(body), r.ContentLength)
	if err != nil {
		log.Printf("%s\nreceived: %+v", err, body)
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	for _, f := range zr.File {
		info := f.FileHeader.FileInfo()
		if !info.IsDir() {
			// do something different if get a zip file here?
			basename := filepath.Base(info.Name())
			name := nameWithoutExt(basename)
			fmt.Fprintf(w, "%s\n", name)
		}
	}
}

func nameWithoutExt(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[:i]
		}
	}
	return filename // no extension; default to input
}
