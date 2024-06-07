package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	superheroes := []string{
		"Batman",
		"Superman",
		"Flash",
	}

	var empty []string

	_ = superheroes

	tpl08 := `
Best Superheroes
{{- range $i, $course := .}}
{{ $i }} {{ $course }}
{{ else }}
There are no elements
{{- end }}
`

	t, err := template.New("exercise_8").Parse(tpl08)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	err = t.Execute(os.Stdout, empty)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
