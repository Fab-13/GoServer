package main

import (
	"fmt"
	"log"
	"net/http"
)

const tpl = `
	RequestURI: %v
	Host:       %v
	Form:       %v
	Some:       %v
`

func reflectForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, fmt.Sprintf(tpl, 
		r.RequestURI, 
		r.Host, 
		r.Form,
		r.Form.Get("Some"),
	))
}

func main() {
	http.HandleFunc("/", reflectForm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}