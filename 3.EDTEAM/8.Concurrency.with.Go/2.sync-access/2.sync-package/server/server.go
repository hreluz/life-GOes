package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {

	duration, _ := strconv.Atoi(r.URL.Query().Get("duration"))
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Fprintf(w, "Duration was %d seconds", duration)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":1234", nil)
}
