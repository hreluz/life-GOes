package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl07 := `
Conditionals
{{- if ge . 18 }}
Welcome to the bar !! Enjoy
{{ else if ge . 12 }}
Sorry, try again next year
{{- else }}
You are too young
{{ end }}
`
	t, err := template.New("exercise_7").Parse(tpl07)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	err = t.Execute(os.Stdout, 14)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
