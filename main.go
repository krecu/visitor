package main

import (
	"net/http"
	"visitor/api"
	"runtime"
	"fmt"
	"visitor/conf"
)

//
func main() {

	conf := config.New()

	fmt.Println(conf)

	runtime.GOMAXPROCS(conf.Cpu)

	apiHttp := api.Method{}

	http.HandleFunc("/api/visitor", apiHttp.Post)
	http.ListenAndServe(conf.Listen, nil)
}