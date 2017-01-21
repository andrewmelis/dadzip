package dadzip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	_ "path/filepath"
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
		fmt.Fprintf(w, "%s\n", f.FileHeader.Name)

		// // zipWalkFunc(f
		// filepath.Walk(f.FileHeader.Name, func(path string, _ os.FileInfo, _ error) error {
		// 	fmt.Fprintf(w, "%s\n", path)
		// 	// log.Printf(path)
		// 	// log.Printf(info.Name())
		// 	return nil
		// })

	}
	// fmt.Fprintf(w, "foobar\n")
}

// func zipWalk(f zip.File, walkFn zipWalkFunc) error {
// 	info, err := f.FileHeader.FileInfo()
// 	if err != nil {
// 		return walkFn(root, nil, err)
// 	}
// 	return walk(root, info, walkFn)
// }

// type zipWalkFunc func(f zip.File, info os.FileInfo, err error) error
