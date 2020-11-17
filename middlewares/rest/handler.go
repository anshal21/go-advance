package rest

import (
	"fmt"
	"net/http"

	"math/rand"

	"github.com/anshal21/go-advance/middlewares/middlewares"
)

func RegisterHandlers() {
	http.HandleFunc("/", middlewares.WithTimer(HelloWorldHandler))

	http.HandleFunc("/random", middlewares.WithAuthorization(map[string]struct{}{
		"admin": {},
	}, GetRandomNumber))

	http.ListenAndServe(":8080", nil)
}

func HelloWorldHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
}

func GetRandomNumber(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(fmt.Sprintf("%v", rand.Int())))
}
