package main

import (
	"net/http"
	"visitor/api"
	"runtime"
	"visitor/conf"
	logger "visitor/log"
)

// основная программа
func main() {

	conf := config.New()
	runtime.GOMAXPROCS(conf.Cpu)

	// стартуем вебсервер
	apiHttp := api.Method{}
	http.HandleFunc("/api/visitor", apiHttp.Post)

	err := http.ListenAndServe(conf.Listen, nil)

	if err != nil {

		logger.Notify(logger.Message{
			ShortMessage:"Failed start web-server: " + err.Error(),
			State: "error",
		})

	}
}