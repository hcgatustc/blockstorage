// Arcsoft.org/arcface/storage
// Storage is a local light smallfile storage system
package main

import (
	"log"
	"net/http"
	"blockstorage/httphandler"
	
)

func main() {
	http.HandleFunc("/arcface-images/", httphandler.Handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

