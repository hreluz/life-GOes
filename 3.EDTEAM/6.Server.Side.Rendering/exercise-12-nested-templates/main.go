package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	t, err := template.New("").ParseGlob("templates/*.tpl")

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	p := person{"Batman", 121}

	err = t.ExecuteTemplate(os.Stdout, "wrapper", &p)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
