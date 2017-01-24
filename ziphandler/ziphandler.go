package ziphandler

import (
	"archive/zip"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func MultiPartZipHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Printf("%s: error parsing form\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, files := range r.MultipartForm.File {
		for _, file := range files {
			log.Printf("opening %s\n", file.Filename)
			f, err := file.Open()
			if err != nil {
				log.Printf("%s: error opening formfile %s\n", err, file.Filename)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer f.Close()

			zr, err := zip.NewReader(f, r.ContentLength)
			if err != nil {
				log.Printf(err.Error())
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}

			for _, zf := range zr.File {
				info := zf.FileHeader.FileInfo()
				if !info.IsDir() {
					// TODO handle nested zip files
					basename := filepath.Base(info.Name())
					name := nameWithoutExt(basename)
					log.Printf("found %s\n", basename)
					fmt.Fprintf(w, "%s\n", name)
				}
			}
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
