package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

const tpl = `
	RequestURI: %v
	Host:       %v
	Form:       %v
	Some:       %v
`

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, fmt.Sprintf(tpl, 
		r.RequestURI, 
		r.Host, 
		r.Form,
		r.Form.Get("Some"),
	))
}

// sends files to the browser
func serveFile(w http.ResponseWriter, r *http.Request) {
	ldir := 
	http.ServeFile(w, r, ldir + r.URL.Path)
}

func main() {
	http.HandleFunc("/form", serveForm)
	http.HandleFunc("/", serveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}