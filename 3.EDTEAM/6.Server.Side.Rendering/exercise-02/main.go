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
	tpl02 := `Hello {{ .Name }}, you are {{ .Age }} years
`
	t, err := template.New("exercise_1").Parse(tpl02)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	p := person{"Batman", 35}
	err = t.Execute(os.Stdout, p)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
