package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	myFunctions := template.FuncMap{
		"usd": func(a float64) string {
			return fmt.Sprintf("$%.fUSD", a)
		},
	}

	var err error
	t, err := template.New("").Funcs(myFunctions).ParseGlob("templates/*.tpl")

	if err != nil {
		log.Fatalf("Error doing parse of the templates: %v", err)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("public/imgs"))))
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("public/data"))))

	http.HandleFunc("/", middleware(t, "index"))
	http.HandleFunc("/courses/", middleware(t, "courses"))

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error when uploading to server: %v", err)
	}
}

func middleware(t *template.Template, funcName string) func(w http.ResponseWriter, r *http.Request) {
	switch funcName {
	case "index":
		return func(w http.ResponseWriter, r *http.Request) {
			index(t, w, r)
		}
	case "courses":
		return func(w http.ResponseWriter, r *http.Request) {
			courses(t, w, r)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}
}

func index(t *template.Template, w http.ResponseWriter, r *http.Request) {
	gp := gridPage{"grid", loadGrid()}
	err := t.ExecuteTemplate(w, "wrapper", gp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func courses(t *template.Template, w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/courses/"):]
	cp := coursePage{"course", getCourse(slug)}
	err := t.ExecuteTemplate(w, "wrapper", &cp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
