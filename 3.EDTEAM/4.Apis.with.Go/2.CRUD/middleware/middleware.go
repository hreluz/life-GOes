package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		log.Printf("request %q, method: %q\n", r.URL.Path, r.Method)
		f(w, r)
		t2 := time.Now()

		fmt.Printf("the request took: %fs\n", t2.Sub(t1).Seconds())
	}
}
