package main

import (
	"log"
	"net/http"

	"github.com/taget/cherry/lissajous"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8123", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	lissajous.lissajous(w)
}
