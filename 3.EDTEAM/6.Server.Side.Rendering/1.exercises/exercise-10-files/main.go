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
	t, err := template.New("invitation.tpl").ParseFiles("./invitation.tpl")

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	p := person{"Batman", 11}

	err = t.Execute(os.Stdout, &p)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
