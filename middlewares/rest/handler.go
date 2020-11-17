package rest

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/", HelloWorldHandler)

	http.ListenAndServe(":8080", nil)
}

func HelloWorldHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
}
