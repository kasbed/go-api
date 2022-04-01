package myroutes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	
	r := mux.NewRouter()

	r.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request) {
		token := DoLogin(r.FormValue("userName"))
		if token != "" {
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, token)
		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "User not allowed.")
		}
		
	}).Methods("POST")

	r.HandleFunc("/user/{id}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		identifier := vars["id"]
		fmt.Fprintf(w, "Wellcome %s", identifier)
	})
	
	r.HandleFunc("/",  func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page")
	})

	return r;
}


func DoLogin(userName string) string {

	return "untoken"
}