package ziphandler

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

var ziptests = []struct {
	testfile string
	expected string
}{
	{"testdata/singlefile.zip", "foo\n"},
	{"testdata/singledir.zip", "foo\n"},
	{"testdata/compositenoextensions.zip", "hello\nbar\nreadonly\n"}, // unclear why these are different. fragile
	{"testdata/compositewithextensions.zip", "bar\nhello\nreadonly\n"},
}

func TestZipHandler(t *testing.T) {
	for _, tt := range ziptests {
		b, err := ioutil.ReadFile(tt.testfile)
		if err != nil {
			t.Errorf("could not open test data")
		}

		r := bytes.NewReader(b)

		req := httptest.NewRequest("POST", "/filenames", r)
		w := httptest.NewRecorder()

		ZipHandler(w, req)

		actual := w.Body.String()

		if tt.expected != actual {
			t.Errorf("expected: %s, actual: %s\n", tt.expected, actual)
		}
	}
}

func TestMultiPartZipHandler(t *testing.T) {
	for _, tt := range ziptests {
		b, err := ioutil.ReadFile(tt.testfile)
		if err != nil {
			t.Errorf("could not open test data")
		}
		r := bytes.NewReader(b)

		body := &bytes.Buffer{}
		mpw := multipart.NewWriter(body)
		part, err := mpw.CreateFormFile("testfile", tt.testfile)
		if err != nil {
			t.Errorf("error creating multipart upload")
		}

		_, err = io.Copy(part, r)
		if err != nil {
			t.Errorf("error copying multipart upload")
		}

		err = mpw.Close()
		if err != nil {
			t.Errorf("error closing multipart writer")
		}

		req := httptest.NewRequest("POST", "/filenames", body)
		req.Header.Set("Content-Type", mpw.FormDataContentType())

		w := httptest.NewRecorder()

		MultiPartZipHandler(w, req)

		actual := w.Body.String()

		if tt.expected != actual {
			t.Errorf("expected: %s, actual: %s\n", tt.expected, actual)
		}
	}
}
