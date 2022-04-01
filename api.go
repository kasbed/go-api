package main

import (
	"net/http"
	"src/myroutes"
)

func main() {
	http.ListenAndServe(":80", myroutes.CreateRoutes())
}