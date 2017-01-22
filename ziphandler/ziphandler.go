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

func MultiPartZipHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received http.Request: %+v\n", r)
	err := r.ParseMultipartForm(10000)
	if err != nil {
		log.Printf("%s: error parsing form\n", err)
		fmt.Fprintf(w, "%s: error parsing form\n", err)
		return
	}

	log.Printf("received http.Request: %+v\n", r)
	log.Printf("received http.Request.MultipartForm: %+v\n", r.MultipartForm)

	// validate that sent file is zip
	// zr, err := zip.NewReader(bytes.NewReader(body), r.ContentLength)
	// if err != nil {
	// 	log.Printf("%s\nreceived: %+v", err, body)
	// 	fmt.Fprintf(w, "%s\n", err)
	// 	return
	// }

	// for _, formFile := range r.MultipartForm.File["testfile"] {
	for k, v := range r.MultipartForm.File {
		log.Printf("\n===================\n")
		log.Printf("%s: %+v\n", k, v)
		for _, formFile := range v {
			file, err := formFile.Open() // return File / ReadCloser
			if err != nil {
				log.Printf("%s: down in the loop\n", err)
				fmt.Fprintf(w, "%s: down in the loop\n", err)
				return
			}
			log.Printf("received f: %+v of type %T\n", formFile.Filename, formFile)
			log.Printf("received file: %+v of type %T\n", file, file)

			zr, err := zip.NewReader(file, r.ContentLength)
			if err != nil {
				log.Printf("%s:\nerror opening zip\n", err)
				fmt.Fprintf(w, "%s\n", err)
				return
			}

			for _, zf := range zr.File {
				info := zf.FileHeader.FileInfo()
				if !info.IsDir() {
					// do something different if get a zip file here?
					basename := filepath.Base(info.Name())
					name := nameWithoutExt(basename)
					fmt.Fprintf(w, "%s\n", name)
				}
			}
		}
	}
}
