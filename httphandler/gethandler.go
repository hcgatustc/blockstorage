package httphandler

import (
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q , Method = %q \n", r.URL.Path, r.Method)
}