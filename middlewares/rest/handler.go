package rest

import (
	"net/http"

	"github.com/anshal21/go-advance/middlewares/middlewares"
)

func RegisterHandlers() {
	http.HandleFunc("/", middlewares.WithTimer(HelloWorldHandler))

	http.ListenAndServe(":8080", nil)
}

func HelloWorldHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
}
