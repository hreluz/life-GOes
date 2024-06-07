package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl05 := `
	Variables, this is how you create a variable in a template
	{{- $myNumber := 41 }}
	The value of the variable is: {{ $myNumber}}
	Reassigning value:
	{{- $myNumber = "hello world" }}
	The new value is: {{ $myNumber }}
`

	t, err := template.New("exercise_5").Parse(tpl05)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	err = t.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
