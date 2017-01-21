package dadzip

import (
	"fmt"
	"net/http"
)

func zipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foobar\n")
	r.Body.String()
}
