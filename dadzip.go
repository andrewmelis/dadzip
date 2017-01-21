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
	// for now, assume

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

		// // zipWalkFunc(f
		// filepath.Walk(f.FileHeader.Name, func(path string, _ os.FileInfo, _ error) error {
		// 	fmt.Fprintf(w, "%s\n", path)
		// 	// log.Printf(path)
		// 	// log.Printf(info.Name())
		// 	return nil
		// })

	}
}
