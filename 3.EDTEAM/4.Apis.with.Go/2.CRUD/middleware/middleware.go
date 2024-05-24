package middleware

import (
	"log"
	"net/http"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request %q, method: %q\n", r.URL.Path, r.Method)
		f(w, r)
	}
}
