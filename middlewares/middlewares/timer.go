package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPHandler func(res http.ResponseWriter, req *http.Request)

func WithTimer(handler HTTPHandler) HTTPHandler {
	return func(res http.ResponseWriter, req *http.Request) {
		stop := startTimer()
		defer stop()
		handler(res, req)
	}
}

func startTimer() func() {
	begin := time.Now()
	return func() {
		fmt.Println("Request Processing Time: ", time.Now().Sub(begin))
	}
}
