package dadzip

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

var ziptests = []struct {
	testfile string
	expected string
}{
	{"testdata/singlefile.zip", "foo\n"},
	{"testdata/singledir.zip", "foo\n"},
	{"testdata/compositenoextensions.zip", "hello\nbar\nreadonly\n"},
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

		zipHandler(w, req)

		actual := w.Body.String()

		if tt.expected != actual {
			t.Errorf("expected: %s, actual: %s\n", tt.expected, actual)
		}
	}
}
