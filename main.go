package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type PublicSuffixList interface {
	PublicSuffixList(domain	string) string
	string() string
}

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	json, _ := json.MarshalIndent(r.Form, "", "  ")
	fmt.Fprintf(w, string(json))
}

type options struct{
	PublicSuffixList PublicSuffixList
	Filename string
	Nopersist bool
}

// sends files to the browser
func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		serveForm(w, r)
		return
	}
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, wd+r.URL.Path)
}

func main() {
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}