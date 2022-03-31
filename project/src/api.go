package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		identifier := vars["id"]

		fmt.Fprintf(w, "Wellcome %s", identifier)
	})
	
	r.HandleFunc("/",  func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page")
	})

	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}