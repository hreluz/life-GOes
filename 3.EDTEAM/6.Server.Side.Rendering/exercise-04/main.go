package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl02 := `Hello 
	{{- . }}
	How are you?
`
	t, err := template.New("exercise_4").Parse(tpl02)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	err = t.Execute(os.Stdout, " Batman")

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
