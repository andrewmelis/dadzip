package dadzip

import (
	"bytes"
	_ "fmt"
	"io/ioutil"
	_ "net/http"
	"net/http/httptest"
	_ "os"
	"testing"
)

func TestZipSingleFile(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/singlefile.zip")
	if err != nil {
		t.Errorf("could not open test data")
	}

	r := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/filenames", r)
	w := httptest.NewRecorder()

	zipHandler(w, req)

	expected := "foo\n"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

func TestZipSingleDir(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/singledir.zip")
	if err != nil {
		t.Errorf("could not open test data")
	}

	r := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/filenames", r)
	w := httptest.NewRecorder()

	zipHandler(w, req)

	expected := "foo\n"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

func TestZipCompositeNoExtensions(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/compositenoextensions.zip")
	if err != nil {
		t.Errorf("could not open test data")
	}

	r := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/filenames", r)
	w := httptest.NewRecorder()

	zipHandler(w, req)

	expected := "hello\nbar\nreadonly\n"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

func TestZipCompositeWithExtensions(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/compositewithextensions.zip")
	if err != nil {
		t.Errorf("could not open test data")
	}

	r := bytes.NewReader(b)

	req := httptest.NewRequest("POST", "/filenames", r)
	w := httptest.NewRecorder()

	zipHandler(w, req)

	expected := "bar\nhello\nreadonly\n"
	actual := w.Body.String()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}
