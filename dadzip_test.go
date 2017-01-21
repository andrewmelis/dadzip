package dadzip

import (
	"archive/zip"
	"bytes"
	_ "fmt"
	"io"
	"log"
	_ "net/http"
	"net/http/httptest"
	"testing"
)

func TestSendZipSingleFile(t *testing.T) {
	testZip := NewTestZip()
	r := httptest.NewRequest("POST", "/filenames", testZip)
	w := httptest.NewRecorder()

	zipHandler(w, r)

	expected := "foo.txt"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

// TODO FIXME
func NewTestZip() io.Reader {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	f, err := w.Create("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("fake file"))
	if err != nil {
		log.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	return buf
}
