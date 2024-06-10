package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func main() {
	myFunctions := template.FuncMap{
		"usd": func(a float64) string {
			return fmt.Sprintf("$%.fUSD", a)
		},
	}

	var err error
	t, err = template.New("").Funcs(myFunctions).ParseGlob("templates/*.tpl")

	if err != nil {
		log.Fatalf("Error doing parse of the templates: %v", err)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("public/imgs"))))
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("public/data"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/courses/", courses)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error when uploading to server: %v", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	gp := gridPage{"grid", loadGrid()}
	err := t.ExecuteTemplate(w, "wrapper", gp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func courses(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/courses/"):]
	cp := coursePage{"course", getCourse(slug)}
	err := t.ExecuteTemplate(w, "wrapper", &cp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
