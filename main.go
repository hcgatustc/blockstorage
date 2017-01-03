// Arcsoft.org/arcface/storage
// Storage is a local light smallfile storage system
package main

import (
	"fmt"
	"log"
	"net/http"
	
)

func main() {
	http.HandleFunc("/arcface-image/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET" :
			
		case "POST" :
		default:
			fmt.Fprintf(w, "URL.Path = %q , Method = %q \n", r.URL.Path,r.Method)
	}
}
