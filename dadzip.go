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
	defer zr.Close()

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
