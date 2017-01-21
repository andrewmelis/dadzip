package dadzip

import (
	"archive/zip"
	"bytes"
	_ "fmt"
	"io"
	"io/ioutil"
	"log"
	_ "net/http"
	"net/http/httptest"
	_ "os"
	"testing"
)

func TestNewSingle(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/singlefile.zip")
	if err != nil {
		t.Errorf("could not open test data")
	}

	r := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/filenames", r)
	w := httptest.NewRecorder()

	zipHandler(w, req)

	expected := "foo.txt\n"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

func TestNewDir(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/singledir.zip")
	if err != nil {
		t.Errorf("could not open test data")
	}

	r := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/filenames", r)
	w := httptest.NewRecorder()

	zipHandler(w, req)

	expected := "foo.txt\n"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}
