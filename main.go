package main

import (
	"net/http"
	"visitor/api"
)

func main() {

	apiHttp := api.Method{}

	http.HandleFunc("/api/visitor", apiHttp.Post)
	http.ListenAndServe(":8080", nil)

}